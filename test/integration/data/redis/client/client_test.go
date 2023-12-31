package client_test

import (
	"testing"

	"github.com/lucky-pocket/luckyPocket-back/test/integration"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	_, closeFunc, err := integration.CreateTestRedisClient()
	defer closeFunc()

	assert.NoError(t, err)
}
