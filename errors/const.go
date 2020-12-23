package errors

import (
	"fmt"
)

type ErrConst string

func (e ErrConst) Error() string {
	return string(e)
}

func NewConst(v string) error {
	return ErrConst(v)
}

func NewfConst(format string, args ...interface{}) error {
	return ErrConst(fmt.Sprintf(format, args...))
}
