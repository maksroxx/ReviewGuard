package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/maksroxx/ReviewGuard/internal/db"
	"github.com/maksroxx/ReviewGuard/internal/models"
	redisclient "github.com/maksroxx/ReviewGuard/internal/redis"
	"github.com/maksroxx/ReviewGuard/internal/service"
	"github.com/redis/go-redis/v9"
)

func ReviewHandler(red redisclient.RedisClient, svc service.HistoryService, rep db.PostgresRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		ip := c.ClientIP()

		var review models.Review
		if err := c.ShouldBindJSON(&review); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		review.ID = uuid.New().String()
		review.UserIP = ip
		review.CreatedAt = time.Now()
		review.Status = "pending"

		if err := rep.SaveReview(ctx, &review); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save review"})
			return
		}

		data, _ := json.Marshal(review)
		if err := red.RDB.XAdd(ctx, &redis.XAddArgs{
			Stream: "moderation_stream",
			Values: map[string]interface{}{"data": data},
		}).Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to enqueue moderation"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "queued_for_moderation"})
	}
}
