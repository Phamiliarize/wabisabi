package application

import (
	"fmt"
	"regexp"

	val "github.com/Phamiliarize/wabisabi/pkg/application/validation"
	"github.com/Phamiliarize/wabisabi/pkg/util"
	"github.com/google/uuid"
)

type VadliationErrors struct {
	Details []val.ValidationDetail
}

func (v VadliationErrors) Error() string {
	return fmt.Sprintf("%v", v.Details)
}

type Session struct {
	TokenID uuid.NullUUID
	UserID  util.NullString
}

func (s *Session) Validate(userIdRegExp *regexp.Regexp) error {
	errors := []val.ValidationDetail{}

	val.ValidateNullUUID([]string{"TokenID"}, s.TokenID, &errors)
	val.ValidateUserID([]string{"UserID"}, s.UserID, userIdRegExp, &errors)

	if len(errors) == 0 {
		return nil
	}

	return VadliationErrors{errors}
}

type Token struct {
	TokenID uuid.NullUUID
}

func (t *Token) Validate() error {
	errors := []val.ValidationDetail{}

	val.ValidateNullUUID([]string{"TokenID"}, t.TokenID, &errors)

	if len(errors) == 0 {
		return nil
	}

	return VadliationErrors{errors}
}

type User struct {
	UserID util.NullString
}

func (u *User) Validate(userIdRegExp *regexp.Regexp) error {
	errors := []val.ValidationDetail{}

	val.ValidateUserID([]string{"UserID"}, u.UserID, userIdRegExp, &errors)

	if len(errors) == 0 {
		return nil
	}

	return VadliationErrors{errors}
}
