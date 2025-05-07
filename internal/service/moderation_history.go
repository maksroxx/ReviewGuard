package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/maksroxx/ReviewGuard/internal/models"
	redisclient "github.com/maksroxx/ReviewGuard/internal/redis"
)

type HistoryService struct {
	redisclient redisclient.RedisClient
}

func NewHistoryService(redis redisclient.RedisClient) *HistoryService {
	return &HistoryService{redis}
}

func (h *HistoryService) SaveToHistory(ctx context.Context, r models.Review) {
	data, _ := json.Marshal(r)
	key := fmt.Sprintf("moderation_history:%s", r.UserIP)
	h.redisclient.RDB.LPush(ctx, key, data)
}

func (h *HistoryService) GetHistory(ctx context.Context, ip string) []models.Review {
	key := fmt.Sprintf("moderation_history:%s", ip)
	data, _ := h.redisclient.RDB.LRange(ctx, key, 0, -1).Result()
	var result []models.Review
	for _, item := range data {
		var r models.Review
		_ = json.Unmarshal([]byte(item), &r)
		result = append(result, r)
	}
	return result
}
