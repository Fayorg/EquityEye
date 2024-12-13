package cache

import (
	"EquityEye/internal/logs"
	"EquityEye/types"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
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

// RegisterProvider This function is not used but it might be useful for other implementation
func (rc *RedisCache) RegisterProvider(provider types.ProviderConfiguration) error {
	return nil
}

func (rc *RedisCache) GetUsage(provider types.ProviderConfiguration) (int, error) {
	length, err := rc.client.XLen(rc.ctx, provider.Name).Result()
	if err != nil {
		return 0, err
	}
	return int(length), nil
}

func (rc *RedisCache) IncreaseUsage(provider types.ProviderConfiguration) error {
	_, err := rc.client.XAdd(rc.ctx, &redis.XAddArgs{
		Stream: provider.Name,
		MinID:  fmt.Sprintf("%d-0", (time.Now().Unix()-int64(provider.LimitTimeframe))*1000),
		Values: map[string]interface{}{
			"timestamp": time.Now().Unix(),
		},
	}).Result()

	return err
}

func (rc *RedisCache) GetProvider(providerType string) (string, error) {
	return "", nil
}
