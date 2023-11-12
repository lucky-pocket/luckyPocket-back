package client

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent"
	"github.com/pkg/errors"
)

// NewClient creates new client from driver and dataSource.
// You should call close to close client.
func NewClient(driver, dataSource string) (client *ent.Client, closeFunc func(), err error) {
	client, err = ent.Open(driver, dataSource)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to open database")
	}

	closeFunc = func() { client.Close() }
	return client, closeFunc, nil
}

// Migrate creates database's schema with given client's schema.
func Migrate(ctx context.Context, client *ent.Client) error {
	if err := client.Schema.Create(ctx); err != nil {
		return errors.Wrap(err, "failed to create schema")
	}
	return nil
}
