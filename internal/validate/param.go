package validate

import (
	"fmt"
	"reflect"

	"github.com/anfilat/go-ews/internal/errors"
	"github.com/anfilat/go-ews/internal/utils"
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
	if utils.IsNil(param) {
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
