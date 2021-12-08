package errors

import (
	"fmt"

	"github.com/anfilat/go-ews/ewsError"
)

func NewSerializationError(val interface{}, localName string) error {
	return &serializationError{val, localName}
}

type serializationError struct {
	val       interface{}
	localName string
}

func (e *serializationError) Error() string {
	return fmt.Sprintf("value %v can't be used for the %s attribute", e.val, e.localName)
}

func (e *serializationError) Is(target error) bool {
	//nolint:serializationError
	return target == ewsError.Serialization
}
