package api

import (
	"go-backend-test/pkg/config"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {

	config, err := config.LoadConfig("./")

	require.NoError(t, err)
	require.NotEmpty(t, config)

	server, err := NewServer(&config)

	require.NoError(t, err)
	require.NotNil(t, server)
}
