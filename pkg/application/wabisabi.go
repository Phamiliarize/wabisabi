package application

import (
	"errors"
	"regexp"

	appError "github.com/Phamiliarize/wabisabi/pkg/application/error"
	val "github.com/Phamiliarize/wabisabi/pkg/application/validation"
)

// Application is an interface describing wabisabi used as a dependency to any interfaces such as
// HTTP or gRPC. It makes it easy to develope new interfaces and worry only about serialization
// concerns.
type Application interface {
	CreateSession(request Session) (string, error)
	ValidateSession(request Token) (bool, error)
	DeleteSessionByTokenId(request Token) error
	DeleteSessionByUserId(request User) error
}

type wabisabi struct {
	datastore    datastore
	userIdRegExp *regexp.Regexp
}

func NewWabisabi(datastore datastore, userIdRegExp *regexp.Regexp) *wabisabi {
	return &wabisabi{
		datastore:    datastore,
		userIdRegExp: userIdRegExp,
	}
}

func (w *wabisabi) CreateSession(request Session) (string, error) {
	err := request.Validate(w.userIdRegExp)
	if err != nil {
		return "", err
	}

	// TODO: need to handle JWT token stuff here :)

	err = w.datastore.CreateSession(request)
	if err != nil {
		if errors.Is(err, appError.ErrSessionCollision) {
			return "", VadliationErrors{
				Details: []val.ValidationDetail{{
					Location: []string{"TokenID"},
					Message:  "TokenID must be unique",
					Value:    request.TokenID,
				}},
			}
		}

		return "", err
	}

	return "", nil
}

func (w *wabisabi) ValidateSession(request Token) (bool, error) {
	err := request.Validate()
	if err != nil {
		return false, err
	}

	_, err = w.datastore.ValidateSession(request)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (w *wabisabi) DeleteSessionByTokenId(request Token) error {
	err := request.Validate()
	if err != nil {
		return err
	}

	err = w.datastore.DeleteSessionByTokenId(request)
	if err != nil {
		return err
	}

	return nil

}

func (w *wabisabi) DeleteSessionByUserId(request User) error {
	err := request.Validate(w.userIdRegExp)
	if err != nil {
		return err
	}

	err = w.datastore.DeleteSessionByUserId(request)
	if err != nil {
		return err
	}

	return nil
}
