package validate

import (
	"fmt"
	"reflect"

	"github.com/anfilat/go-ews/internal/errors"
)

type selfValidate interface {
	Validate() error
}

func Param(param interface{}, paramName string) error {
	switch val := param.(type) {
	case string:
		if param == "" {
			return errors.NewValidateError(fmt.Sprintf("argument %s is empty", paramName))
		}
	case selfValidate:
		if err := val.Validate(); err != nil {
			return errors.NewValidateErrorWithWrap(err, fmt.Sprintf("invalid parameter %s: %s", paramName, err))
		}
	default:
		return errors.NewValidateError(fmt.Sprintf("invalid param type for %s", paramName))
	}
	return nil
}

func ParamSlice(param interface{}, paramName string) error {
	slice := reflect.ValueOf(param)
	if slice.Kind() != reflect.Slice || slice.IsNil() || slice.Len() == 0 {
		return errors.NewValidateError(fmt.Sprintf("the collection %s is empty", paramName))
	}

	for i := 0; i < slice.Len(); i++ {
		if err := Param(slice.Index(i).Interface(), fmt.Sprintf("%s [%d]", paramName, i)); err != nil {
			return err
		}
	}
	return nil
}

func ParamAllowNull(param interface{}, paramName string) error {
	if isNil(param) {
		return nil
	}

	val, ok := param.(selfValidate)
	if !ok {
		return nil
	}

	if err := val.Validate(); err != nil {
		return errors.NewValidateErrorWithWrap(err, fmt.Sprintf("invalid parameter %s: %s", paramName, err))
	}

	return nil
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	//nolint:exhaustive
	switch reflect.TypeOf(i).Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	default:
		return false
	}
}
