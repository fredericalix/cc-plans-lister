package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	token := "test_token_123"
	client := NewClient(token)
	
	require.NotNil(t, client)
	require.NotNil(t, client.cc)
}

// Note: Integration tests for GetAddonProviders and GetProductInstances
// would require a valid API token and network access, so they are not
// included in unit tests. These should be tested separately as integration tests.