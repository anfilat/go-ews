package ewsError

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	err := NewValidateError("oops")
	require.ErrorIs(t, err, Validate)

	err = fmt.Errorf("an error: %w", err)
	require.ErrorIs(t, err, Validate)
}

func TestValidateWithWrap(t *testing.T) {
	err := NewValidateErrorWithWrap(io.EOF, "an error")
	require.ErrorIs(t, err, io.EOF)
	require.ErrorIs(t, err, Validate)
}
