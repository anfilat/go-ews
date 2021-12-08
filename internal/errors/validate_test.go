package errors

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/anfilat/go-ews/ewsError"
)

func TestValidate(t *testing.T) {
	err := NewValidateError("oops")
	require.ErrorIs(t, err, ewsError.Validate)

	err = fmt.Errorf("an error: %w", err)
	require.ErrorIs(t, err, ewsError.Validate)
}

func TestValidateWithWrap(t *testing.T) {
	err := NewValidateErrorWithWrap(io.EOF, "an error")
	require.ErrorIs(t, err, io.EOF)
	require.ErrorIs(t, err, ewsError.Validate)
}
