package cache

import (
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	rdb *redis.Client
}

func New(
	rdb *redis.Client,
) *RedisCache {

	return &RedisCache{
		rdb: rdb,
	}
}

func (c *RedisCache) Get(
	ctx context.Context,
	key string,
) (bool, error) {

	value, err := c.rdb.Get(
		ctx,
		key,
	).Result()

	fmt.Println("========== REDIS GET ==========")
	fmt.Println("KEY   :", key)
	fmt.Println("VALUE :", value)
	fmt.Println("ERROR :", err)
	fmt.Println("===============================")

	if err != nil {
		return false, err
	}

	return strconv.ParseBool(
		value,
	)
}

func (c *RedisCache) Set(
	ctx context.Context,
	key string,
	enabled bool,
) error {

	err := c.rdb.Set(
		ctx,
		key,
		strconv.FormatBool(enabled),
		0,
	).Err()

	fmt.Println("========== REDIS SET ==========")
	fmt.Println("KEY   :", key)
	fmt.Println("VALUE :", enabled)
	fmt.Println("ERROR :", err)
	fmt.Println("===============================")

	return err
}

func (c *RedisCache) Delete(
	ctx context.Context,
	key string,
) error {

	err := c.rdb.Del(
		ctx,
		key,
	).Err()

	fmt.Println("======== REDIS DELETE =========")
	fmt.Println("KEY   :", key)
	fmt.Println("ERROR :", err)
	fmt.Println("===============================")

	return err
}
