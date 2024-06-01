package error

import (
	"errors"
)

var (
	ErrSessionCollision = errors.New(("collision detected on token id"))
)
