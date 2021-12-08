package errors

import "github.com/anfilat/go-ews/ewsError"

func NewValidateError(text string) error {
	return &validateError{nil, text}
}

func NewValidateErrorWithWrap(err error, text string) error {
	return &validateError{err, text}
}

type validateError struct {
	err error
	s   string
}

func (e *validateError) Error() string {
	return e.s
}

func (e *validateError) Unwrap() error {
	return e.err
}

func (e *validateError) Is(target error) bool {
	//nolint:errorlint
	return target == ewsError.Validate
}
