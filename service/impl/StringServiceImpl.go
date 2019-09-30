package impl

import (
	"errors"
	"strings"
)

// StringServiceImpl is an implementation of the StringService
type StringServiceImpl struct{}

//Uppercase service
func (StringServiceImpl) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

//Count service
func (StringServiceImpl) Count(s string) int {
	return len(s)
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty string")
