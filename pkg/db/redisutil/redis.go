package redisutil

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	ClusterMode bool     `mapstructure:"clusterMode"` // Whether to use Redis in cluster mode.
	Address     []string `mapstructure:"address"`     // List of Redis server addresses (host:port).
	Username    string   `mapstructure:"username"`    // Username for Redis authentication (Redis 6 ACL).
	Password    string   `mapstructure:"password"`    // Password for Redis authentication.
	MaxRetry    int      `mapstructure:"maxRetry"`    // Maximum number of retries for a command.
	DB          int      `mapstructure:"db"`          // Database number to connect to, for non-cluster mode.
	PoolSize    int      `mapstructure:"poolSize"`    // Number of connections to pool.
}

func NewRedisClient(ctx context.Context, config *Config) (redis.UniversalClient, error) {
	if len(config.Address) == 0 {
		return nil, errors.New("redis address is empty")
	}
	var cli redis.UniversalClient
	if config.ClusterMode || len(config.Address) > 1 {
		opt := &redis.ClusterOptions{
			Addrs:      config.Address,
			Username:   config.Username,
			Password:   config.Password,
			PoolSize:   config.PoolSize,
			MaxRetries: config.MaxRetry,
		}
		cli = redis.NewClusterClient(opt)
	} else {
		opt := &redis.Options{
			Addr:       config.Address[0],
			Username:   config.Username,
			Password:   config.Password,
			DB:         config.DB,
			PoolSize:   config.PoolSize,
			MaxRetries: config.MaxRetry,
		}
		cli = redis.NewClient(opt)
	}
	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, errors.New(fmt.Sprintf("[ERROR]: %s, Redis Ping failed, Address: %s, Username: %s, ClusterMode: %v", err, config.Address, config.Username, config.ClusterMode))
	}
	return cli, nil
}
