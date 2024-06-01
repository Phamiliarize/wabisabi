package sqlite

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Phamiliarize/wabisabi/pkg/application"
	appError "github.com/Phamiliarize/wabisabi/pkg/application/error"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type sqlite struct {
	db *sql.DB
}

func NewSqLite() *sqlite {
	var database sqlite
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared&_journal_mode=WAL")
	if err != nil {
		panic(fmt.Sprintf("failed to open sqlite3 DB: %v", err))
	}
	database.db = db

	err = database.setupTable()
	if err != nil {
		panic(fmt.Sprintf("failed to open sqlite3 DB: %v", err))
	}

	return &database
}

func (s *sqlite) setupTable() error {
	_, err := s.db.Exec(`
	CREATE TABLE session(token_id blob NOT NULL PRIMARY KEY, user_id blob NOT NULL);
	CREATE INDEX user_id_idx ON session(user_id);
	`)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(`INSERT INTO session (token_id, user_id) VALUES (?, ?);`, []byte(uuid.NewString()), []byte(uuid.NewString()))
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlite) CreateSession(req application.Session) error {
	_, err := s.db.Exec(
		`INSERT INTO session (token_id, user_id) VALUES (?, ?);`,
		[]byte(req.TokenID.UUID.String()),
		req.UserID.String,
	)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "unique") {
			return fmt.Errorf(
				"non-unique token id: %s, %w",
				req.TokenID.UUID.String(),
				appError.ErrSessionCollision,
			)
		}
		return err
	}
	return nil
}

func (s *sqlite) ValidateSession(req application.Token) (bool, error) {
	token := Session{}
	row := s.db.QueryRow(
		`SELECT * FROM session WHERE token_id = ?;`,
		[]byte(req.TokenID.UUID.String()),
	)

	err := row.Scan(&token.TokenID, &token.UserID)
	if err != nil {
		return false, fmt.Errorf("invalid_token")
	}

	return true, nil
}

func (s *sqlite) DeleteSessionByTokenId(req application.Token) error {
	_, err := s.db.Exec(
		`DELETE FROM session where token_id = ?;`,
		[]byte(req.TokenID.UUID.String()),
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlite) DeleteSessionByUserId(req application.User) error {
	_, err := s.db.Exec(
		`DELETE FROM session where user_id = ?;`,
		req.UserID.String,
	)
	if err != nil {
		return err
	}

	return nil
}
