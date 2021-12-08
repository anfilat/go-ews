package validate

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/anfilat/go-ews/ewsError"
	"github.com/anfilat/go-ews/internal/errors"
)

func TestValidateFilled(t *testing.T) {
	err := Param("string", "name")
	require.NoError(t, err)

	err = Param(validateData{data: 42}, "name")
	require.NoError(t, err)

	err = ParamSlice([]validateData{{data: 42}, {data: 7}}, "name")
	require.NoError(t, err)
}

func TestValidateEmpty(t *testing.T) {
	err := Param("", "name")
	require.ErrorIs(t, err, ewsError.Validate)

	err = Param(validateData{data: 0}, "name")
	require.ErrorIs(t, err, ewsError.Validate)

	err = ParamSlice(nil, "name")
	require.ErrorIs(t, err, ewsError.Validate)

	err = ParamSlice([]validateData{}, "name")
	require.ErrorIs(t, err, ewsError.Validate)

	err = ParamSlice([]validateData{{data: 42}, {data: 0}}, "name")
	require.ErrorIs(t, err, ewsError.Validate)
}

type validateData struct {
	data int
}

func (v validateData) Validate() error {
	if v.data == 0 {
		return errors.NewValidateError("data is empty")
	}
	return nil
}
