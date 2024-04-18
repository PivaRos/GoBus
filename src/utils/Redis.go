package utils

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func CheckRedisConnection(ctx context.Context, rdb *redis.Client) error {
	// Send a PING command to Redis
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
