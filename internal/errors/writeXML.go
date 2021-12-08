package errors

import "github.com/anfilat/go-ews/ewsError"

func NewWriteXMLError(err error) error {
	return &writeXMLError{err}
}

type writeXMLError struct {
	err error
}

func (e *writeXMLError) Error() string {
	return e.err.Error()
}

func (e *writeXMLError) Unwrap() error {
	return e.err
}

func (e *writeXMLError) Is(target error) bool {
	//nolint:errorlint
	return target == ewsError.WriteXML
}
