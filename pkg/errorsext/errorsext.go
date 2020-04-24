package errorsext

import (
	"errors"
	"strings"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrConflict     = errors.New("conflict")
	ErrUnauthorized = errors.New("unauthorized")
	ErrValidation   = errors.New("validation")
)

func IsNotFound(err error) bool {
	return strings.HasSuffix(err.Error(), ErrNotFound.Error())
}
func IsConflict(err error) bool {
	return strings.HasSuffix(err.Error(), ErrConflict.Error())
}
func IsUnauthorized(err error) bool {
	return strings.HasSuffix(err.Error(), ErrUnauthorized.Error())
}
func IsValidation(err error) bool {
	return strings.HasSuffix(err.Error(), ErrValidation.Error())
}
