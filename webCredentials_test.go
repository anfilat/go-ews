package ews

import (
	"context"
	"encoding/base64"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWebCredentialsPrepareWebRequest(t *testing.T) {
	const userName = "userName"
	const password = "password"

	credentials := NewWebCredentials(userName, password)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "", nil)
	require.NoError(t, err)

	credentials.PrepareWebRequest(req)
	auth, ok := req.Header["Authorization"]
	require.True(t, ok)
	require.Len(t, auth, 1)
	expectAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(userName+":"+password))
	require.Equal(t, expectAuth, auth[0])
}
