package config

import (
	"github.com/redis/go-redis/v9"
	"sync"
)

type RedisConfigMode int

const (
	RedisConfigModeOptions RedisConfigMode = 0
	RedisConfigModeRaw     RedisConfigMode = 1
)

type RedisConfig struct {
	Mode RedisConfigMode

	Options *redis.UniversalOptions
	Client  redis.UniversalClient

	once sync.Once
}

func (c *RedisConfig) InitClient() redis.UniversalClient {
	c.once.Do(func() {
		if c.Mode == RedisConfigModeOptions {
			c.Client = redis.NewUniversalClient(c.Options)
		}
	})
	return c.Client
}
