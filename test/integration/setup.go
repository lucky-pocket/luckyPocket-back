package integration

import (
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/client"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent"
	"github.com/pkg/errors"
)

func CreateTestClient() (c *ent.Client, closeFunc func(), err error) {
	driver, dataSource := client.NewMemorySQLiteDialect()
	c, closeFunc, err = client.NewClient(driver, dataSource)
	c = c.Debug()
	if err != nil {
		return nil, nil, errors.Wrap(err, "error createing test client")
	}
	return
}
