package integration

import (
	ent_client "github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/client"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent"
	redis_client "github.com/lucky-pocket/luckyPocket-back/internal/infra/data/redis/client"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

func CreateTestEntClient() (c *ent.Client, closeFunc func(), err error) {
	driver, dataSource := ent_client.NewMemorySQLiteDialect()
	c, closeFunc, err = ent_client.NewClient(driver, dataSource)
	c = c.Debug()
	if err != nil {
		return nil, nil, errors.Wrap(err, "error createing test client")
	}
	return
}

func CreateTestRedisClient() (c *redis.Client, closeFunc func(), err error) {
	c, closeFunc, err = redis_client.NewClient("localhost:6379", "", 0)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error createing test client")
	}
	return
}
