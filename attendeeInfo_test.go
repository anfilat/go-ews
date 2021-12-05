package ews

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAttendeeInfoValidateSuccess(t *testing.T) {
	err := NewAttendeeInfo("mail@mail.com").Validate()
	require.NoError(t, err)
}

func TestAttendeeInfoValidateFail(t *testing.T) {
	err := NewAttendeeInfo("").Validate()
	require.Error(t, err)
}
