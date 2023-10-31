package client_test

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/client"
	"github.com/lucky-pocket/luckyPocket-back/test/integration"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c, closeFunc, err := integration.CreateTestEntClient()

	if assert.NoError(t, err) {
		defer closeFunc()

		err := client.Migrate(context.Background(), c)
		assert.NoError(t, err)
	}
}
