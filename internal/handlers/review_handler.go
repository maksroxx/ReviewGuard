package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/maksroxx/ReviewGuard/internal/db"
	"github.com/maksroxx/ReviewGuard/internal/models"
	"github.com/maksroxx/ReviewGuard/internal/redis"
	"github.com/maksroxx/ReviewGuard/internal/service"
)

func ReviewHandler(redis redis.RedisClient, svc service.HistoryService, rep db.PostgresRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		ip := c.ClientIP()

		// limitMinute := redis_rate.PerMinute(5)
		// limitHour := redis_rate.Limit{Rate: 20, Period: time.Hour}

		// resMin, err := redis.Limiter.Allow(ctx, ip+":minute", limitMinute)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "rate limit error"})
		// 	return
		// }
		// if resMin.Allowed == 0 {
		// 	c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests (minute limit)", "retry_after": resMin.RetryAfter.Seconds()})
		// 	return
		// }

		// resHour, err := redis.Limiter.Allow(ctx, ip+":hour", limitHour)
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "rate limit error"})
		// 	return
		// }
		// if resHour.Allowed == 0 {
		// 	c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests (hour limit)", "retry_after": resHour.RetryAfter.Seconds()})
		// 	return
		// }

		var review models.Review
		if err := c.ShouldBindJSON(&review); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		review.ID = uuid.New().String()
		review.UserIP = ip
		review.CreatedAt = time.Now()

		if service.ContainsBannedWords(review.Content) || service.ContainsLinks(review.Content) {
			review.Status = "moderation"
		} else {
			review.Status = "approved"
		}

		svc.SaveToHistory(ctx, review)
		rep.SaveReview(ctx, &review)

		c.JSON(http.StatusOK, gin.H{"status": review.Status})
	}
}
