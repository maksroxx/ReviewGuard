package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/maksroxx/ReviewGuard/internal/models"
	redisclient "github.com/maksroxx/ReviewGuard/internal/redis"
)

type CacheSerice struct {
	redisclient redisclient.RedisClient
}

func NewCacheService(redis redisclient.RedisClient) *CacheSerice {
	return &CacheSerice{redis}
}

func (c *CacheSerice) CacheIPStats(ctx context.Context, period string, stats []models.IPStats) error {
	data, err := json.Marshal(stats)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("ip_stats:%s", period)
	return c.redisclient.RDB.Set(ctx, key, data, time.Minute*10).Err()
}

func (c *CacheSerice) GetCachedIPStats(ctx context.Context, period string) ([]models.IPStats, error) {
	key := fmt.Sprintf("ip_stats:%s", period)
	data, err := c.redisclient.RDB.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var stats []models.IPStats
	if err := json.Unmarshal([]byte(data), &stats); err != nil {
		return nil, err
	}
	return stats, nil
}
