package rest

import (
	"github.com/Phamiliarize/wabisabi/pkg/util"
	"github.com/google/uuid"
)

type GenericErrorMessageJSON struct {
	Message string `json:"msg,omitempty"`
}

type CreateSessionJSON struct {
	TokenID uuid.NullUUID   `json:"TokenID"`
	UserID  util.NullString `json:"UserID"`
}

type ValidateSessionJSON struct {
	TokenID uuid.NullUUID `json:"TokenID"`
}

type DeleteSessionByTokenJSON struct {
	TokenID uuid.NullUUID `json:"TokenID"`
}

type DeleteSessionByUserJSON struct {
	UserID util.NullString `json:"UserID"`
}
