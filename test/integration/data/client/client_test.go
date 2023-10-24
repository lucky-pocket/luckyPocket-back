package client_test

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"

	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/client"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent"
	"github.com/stretchr/testify/assert"
)

func CreateTestClient() (c *ent.Client, closeFunc func(), err error) {
	driver, dataSource := client.NewMemorySQLiteDialect()
	c, closeFunc, err = client.NewClient(driver, dataSource)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error createing test client")
	}
	return
}

func TestNewClient(t *testing.T) {
	c, closeFunc, err := CreateTestClient()

	if assert.NoError(t, err) {
		defer closeFunc()

		err := client.Migrate(context.Background(), c)
		assert.NoError(t, err)
	}
}
