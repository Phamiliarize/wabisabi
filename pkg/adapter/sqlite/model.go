package sqlite

import (
	"database/sql"

	"github.com/google/uuid"
)

type Session struct {
	TokenID uuid.UUID      `sql:"token_id"`
	UserID  sql.NullString `sql:"user_id"`
}
