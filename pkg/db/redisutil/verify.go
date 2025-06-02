package redisutil

import (
	"context"
	"errors"
	"fmt"
)

// Check CheckRedis checks the Redis connection.
func Check(ctx context.Context, config *Config) error {
	client, err := NewRedisClient(ctx, config)
	if err != nil {
		return err
	}
	defer client.Close()

	// Ping the Redis server to check connectivity.
	if err := client.Ping(ctx).Err(); err != nil {
		return errors.New(fmt.Sprintf("[ERROR]: %s, Redis ping failed, config: %v", err, config))
	}

	return nil
}
