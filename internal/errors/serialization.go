package errors

import (
	"fmt"

	"github.com/anfilat/go-ews/ewsError"
)

func NewAttrSerializationError(val interface{}, localName string) error {
	return &attrSerializationError{val, localName}
}

type attrSerializationError struct {
	val       interface{}
	localName string
}

func (e *attrSerializationError) Error() string {
	return fmt.Sprintf("value %v can't be used for the %s attribute", e.val, e.localName)
}

func (e *attrSerializationError) Is(target error) bool {
	//nolint:errorlint
	return target == ewsError.Serialization
}

func NewValueSerializationError(val interface{}, localName string) error {
	return &valueSerializationError{val, localName}
}

type valueSerializationError struct {
	val       interface{}
	localName string
}

func (e *valueSerializationError) Error() string {
	return fmt.Sprintf("values %v can't be used for the %s element", e.val, e.localName)
}

func (e *valueSerializationError) Is(target error) bool {
	//nolint:errorlint
	return target == ewsError.Serialization
}
