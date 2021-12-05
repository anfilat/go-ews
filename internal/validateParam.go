package internal

import (
	"fmt"
	"reflect"
)

func ValidateParam(param interface{}, paramName string) error {
	type SelfValidate interface {
		Validate() error
	}

	switch val := param.(type) {
	case string:
		if param == "" {
			return fmt.Errorf("argument %s is empty", paramName)
		}
	case SelfValidate:
		if err := val.Validate(); err != nil {
			return fmt.Errorf("invalid parameter %s: %w", paramName, err)
		}
	default:
		return fmt.Errorf("invalid param type for %s", paramName)
	}
	return nil
}

func ValidateParamSlice(param interface{}, paramName string) error {
	slice := reflect.ValueOf(param)
	if slice.Kind() != reflect.Slice || slice.IsNil() || slice.Len() == 0 {
		return fmt.Errorf("the collection %s is empty", paramName)
	}

	for i := 0; i < slice.Len(); i++ {
		if err := ValidateParam(slice.Index(i).Interface(), fmt.Sprintf("%s [%d]", paramName, i)); err != nil {
			return err
		}
	}
	return nil
}

func ValidateParamAllowNull(param interface{}, paramName string) error {
	if isNil(param) {
		return nil
	}

	if val, ok := param.(interface{ Validate() error }); ok {
		if err := val.Validate(); err != nil {
			return fmt.Errorf("invalid parameter %s: %w", paramName, err)
		}
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
