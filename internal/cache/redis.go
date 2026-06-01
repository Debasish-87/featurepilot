package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/Debasish-87/featurepilot/internal/config"
)

func NewRedis(cfg *config.Config) *redis.Client {

	addr := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)

	return redis.NewClient(
		&redis.Options{
			Addr: addr,
		},
	)
}

func Ping(rdb *redis.Client) error {
	return rdb.Ping(context.Background()).Err()
}
