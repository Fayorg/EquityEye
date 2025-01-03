package cache

import (
	"EquityEye/internal/logs"
	"EquityEye/types"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
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
	// Get all the data from the stream
	data, err := rc.client.XRange(rc.ctx, provider.Name, "-", "+").Result()
	if err != nil {
		return 0, err
	}

	// Calculate the usage
	usage := 0
	for _, d := range data {
		for k, v := range d.Values {
			if k == "used" {
				val, err := strconv.Atoi(v.(string))
				if err != nil {
					return 0, err
				}
				usage += val
			}
		}
	}

	return usage, nil
}

func (rc *RedisCache) IncreaseUsage(provider types.ProviderConfiguration) error {
	return rc.IncreaseUsageBy(provider, 1)
}

func (rc *RedisCache) IncreaseUsageBy(provider types.ProviderConfiguration, value int) error {
	_, err := rc.client.XAdd(rc.ctx, &redis.XAddArgs{
		Stream: provider.Name,
		MinID:  fmt.Sprintf("%d-0", (time.Now().Unix()-int64(provider.LimitTimeframe))*1000),
		Values: map[string]interface{}{
			"timestamp": time.Now().Unix(),
			"used":      value,
		},
	}).Result()

	return err
}

func (rc *RedisCache) TemporaryDisableProvider(provider types.ProviderConfiguration, duration time.Duration) error {
	_, err := rc.client.Set(rc.ctx, fmt.Sprintf("%s:disabled", provider.Name), time.Now().Add(duration), duration).Result()
	return err
}

func (rc *RedisCache) IsProviderTemporarilyDisabled(provider types.ProviderConfiguration) (bool, error) {
	ttl, err := rc.client.TTL(rc.ctx, fmt.Sprintf("%s:disabled", provider.Name)).Result()
	if err != nil {
		return false, err
	}
	return ttl > 0, nil
}

func (rc *RedisCache) GetProvider(providerType string) (string, error) {
	return "", nil
}
