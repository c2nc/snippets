package errors

import (
	"fmt"
)

func NewConst(v string) error {
	return ErrConst(v)
}

func NewfConst(format string, args ...interface{}) error {
	return ErrConst(fmt.Sprintf(format, args...))
}