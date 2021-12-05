package ews

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTimeWindowValidateSuccess(t *testing.T) {
	now := time.Now()

	err := NewTimeWindow(now, now.Add(time.Second)).Validate()
	require.NoError(t, err)
}

func TestTimeWindowValidateFail(t *testing.T) {
	now := time.Now()

	err := NewTimeWindow(now, now).Validate()
	require.Error(t, err)

	err = NewTimeWindow(now, now.Add(-time.Second)).Validate()
	require.Error(t, err)
}
