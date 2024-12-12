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

func (rc *RedisCache) RegisterProvider(provider types.Provider) error {
	_, err := rc.client.XAdd(rc.ctx, &redis.XAddArgs{
		Stream: provider.Name,
		MinID:  fmt.Sprintf("%d-0", (time.Now().Unix()-int64(provider.LimitTimeframe))*1000),
		Values: map[string]interface{}{
			"timestamp": time.Now().Unix(),
		},
	}).Result()

	return err
}

func (rc *RedisCache) GetUsage(provider types.Provider) (int, error) {
	length, err := rc.client.XLen(rc.ctx, provider.Name).Result()
	if err != nil {
		return 0, err
	}
	return int(length), nil
}

func (rc *RedisCache) IncreaseUsage(provider types.Provider) (int, error) {
	return 0, nil
}

func (rc *RedisCache) GetProvider(providerName []string) types.Provider {
	return types.Provider{}
}
