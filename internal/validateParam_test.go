package internal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateSuccess(t *testing.T) {
	err := ValidateParam("string", "name")
	require.NoError(t, err)

	err = ValidateParam(validateData{data: 42}, "name")
	require.NoError(t, err)

	err = ValidateParamSlice([]validateData{{data: 42}, {data: 7}}, "name")
	require.NoError(t, err)
}

func TestValidateFail(t *testing.T) {
	err := ValidateParam("", "name")
	require.Error(t, err)

	err = ValidateParam(validateData{data: 0}, "name")
	require.Error(t, err)

	err = ValidateParamSlice(nil, "name")
	require.Error(t, err)

	err = ValidateParamSlice([]validateData{}, "name")
	require.Error(t, err)

	err = ValidateParamSlice([]validateData{{data: 42}, {data: 0}}, "name")
	require.Error(t, err)
}

type validateData struct {
	data int
}

func (v validateData) Validate() error {
	if v.data == 0 {
		return fmt.Errorf("data is empty")
	}
	return nil
}
