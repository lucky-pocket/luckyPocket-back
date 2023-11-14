package client

import (
	"context"
	entsql "entgo.io/ent/dialect/sql"

	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent"
	"github.com/pkg/errors"

	"database/sql"
)

// NewClient creates new client from driver and dataSource.
// You should call close to close client.
func NewClient(driver, dataSource string) (client *ent.Client, sqlDB *sql.DB, closeFunc func(), err error) {
	drv, err := entsql.Open(driver, dataSource)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to open database")
	}

	closeFunc = func() { drv.Close() }
	return ent.NewClient(ent.Driver(drv)), drv.DB(), closeFunc, nil
}

// Migrate creates database's schema with given client's schema.
func Migrate(ctx context.Context, client *ent.Client) error {
	if err := client.Schema.Create(ctx); err != nil {
		return errors.Wrap(err, "failed to create schema")
	}
	return nil
}
