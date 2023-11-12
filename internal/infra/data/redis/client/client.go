package client

import (
	"context"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

func NewClient(addr, password string, db int) (c *redis.Client, closeFunc func(), err error) {
	c = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	cmd := c.Ping(context.Background())
	if err := cmd.Err(); err != nil {
		return nil, nil, errors.Wrap(err, "client ping failed")
	}

	closeFunc = func() {
		c.Close()
	}

	return c, closeFunc, nil
}
