package utils

type ErrArgument struct {
	Wrapped error
}

func (e ErrArgument) Error() string {
	return "invalid argument"
}

func (e ErrArgument) Unwrap() error {
	return e.Wrapped
}

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func (e ErrorResponse) Error() string {
	return e.ErrorMessage
}
