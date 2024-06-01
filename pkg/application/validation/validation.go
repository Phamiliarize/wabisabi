package validation

import (
	"fmt"
	"regexp"

	"github.com/Phamiliarize/wabisabi/pkg/util"
	"github.com/google/uuid"
)

func ValidateNullUUID(loc []string, val uuid.NullUUID, destination *[]ValidationDetail) {
	if !val.Valid || val.UUID == uuid.Nil {
		*destination = append(*destination, ValidationDetail{
			Location: loc,
			Message:  "Expected a valid UUID value.",
			Value:    val,
		})
	}
}

func ValidateUserID(loc []string, val util.NullString, regexp *regexp.Regexp, destination *[]ValidationDetail) {
	if !val.Valid || val.String == "" || !regexp.MatchString(val.String) {
		*destination = append(*destination, ValidationDetail{
			Location: loc,
			Message:  fmt.Sprintf("Expected value to match regular expression: %s", regexp.String()),
			Value:    val,
		})
	}
}
