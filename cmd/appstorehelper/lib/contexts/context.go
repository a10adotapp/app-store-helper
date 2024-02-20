package contexts

import (
	"errors"
)

type contextKey struct {
	name string
}

var (
	ErrNoValue = errors.New("no value found in context")
)

func (ck contextKey) String() string {
	return ck.name
}

func NewNoValueError(err error) error {
	return errors.Join(ErrNoValue, err)
}
