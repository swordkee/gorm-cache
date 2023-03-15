package cache

import (
	"fmt"
	"gorm.io/gorm"

	"github.com/redis/go-redis/v9"
	"github.com/swordkee/gorm-cache/config"
)

func NewPlugin(opts ...Option) gorm.Plugin {
	cacheConfig := newCache(opts...)
	cache := &Gorm2Cache{
		Config: cacheConfig,
	}
	err := cache.Init()
	if err != nil {
		return nil
	}
	return cache
}

func newCache(opts ...Option) *config.CacheConfig {
	opt := new(config.CacheConfig)
	for _, f := range opts {
		f(opt)
	}
	if len(opts) == 0 {
		return &config.CacheConfig{
			CacheLevel:           config.CacheLevelAll,
			CacheStorage:         config.CacheStorageMemory,
			InvalidateWhenUpdate: true,
			CacheTTL:             5000,
			CacheMaxItemCnt:      5,
		}
	}
	return &config.CacheConfig{
		CacheLevel:           opt.CacheLevel,
		CacheStorage:         opt.CacheStorage,
		RedisConfig:          opt.RedisConfig,
		Tables:               opt.Tables,
		InvalidateWhenUpdate: opt.InvalidateWhenUpdate,
		CacheTTL:             opt.CacheTTL,
		CacheMaxItemCnt:      opt.CacheMaxItemCnt,
		CacheSize:            opt.CacheSize,
		DebugMode:            opt.DebugMode,
		DebugLogger:          opt.DebugLogger,
	}
}

func NewGorm2Cache(cacheConfig *config.CacheConfig) (*Gorm2Cache, error) {
	if cacheConfig == nil {
		return nil, fmt.Errorf("you pass a nil config")
	}
	cache := &Gorm2Cache{
		Config: cacheConfig,
	}
	err := cache.Init()
	if err != nil {
		return nil, err
	}
	return cache, err
}
func NewRedisConfigWithOptions(options *redis.UniversalOptions) *config.RedisConfig {
	return &config.RedisConfig{
		Mode:    config.RedisConfigModeOptions,
		Options: options,
	}
}

func NewRedisConfigWithClient(client redis.UniversalClient) *config.RedisConfig {
	return &config.RedisConfig{
		Mode:   config.RedisConfigModeRaw,
		Client: client,
	}
}
