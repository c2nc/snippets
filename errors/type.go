package errors

type ErrConst string

func (e ErrConst) Error() string {
	return string(e)
}
