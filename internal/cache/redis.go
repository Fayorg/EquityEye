package cache

import (
	"EquityEye/internal/logs"
	"EquityEye/types"
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisCache(url string) *RedisCache {
	opt, err := redis.ParseURL(url)
	if err != nil {
		logs.Error("Couldn't not parse cache url")
		panic(err)
	}

	rdb := redis.NewClient(opt)

	return &RedisCache{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (rc *RedisCache) RegisterProvider(provider types.Provider) error {
	return nil
}

func (rc *RedisCache) IncreaseUsage(provider types.Provider) (int, error) {
	return 0, nil
}

func (rc *RedisCache) GetProvider(providerName []string) types.Provider {
	return types.Provider{}
}
