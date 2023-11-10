package error

type TError struct {
	Code int
	Err  error
}

func NewError(code int, err error) *TError {
	return &TError{
		Code: code,
		Err:  err,
	}

}
