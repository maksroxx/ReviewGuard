package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
	"github.com/maksroxx/ReviewGuard/internal/models"
	"github.com/maksroxx/ReviewGuard/internal/redis"
	"github.com/maksroxx/ReviewGuard/internal/service"
)

func ReviewHandler(c *gin.Context) {
	ctx := context.Background()
	ip := c.ClientIP()

	limit1 := redis_rate.PerMinute(5)
	limit2 := redis_rate.Limit{Rate: 20, Period: time.Hour}

	res1, err := redis.Limiter.Allow(ctx, ip+":m", limit1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "rate limit error"})
		return
	}
	if res1.Allowed == 0 {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests (minute limit)", "retry_after": res1.RetryAfter.Seconds()})
		return
	}

	res2, err := redis.Limiter.Allow(ctx, ip+":h", limit2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "rate limit error"})
		return
	}
	if res2.Allowed == 0 {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests (hour limit)", "retry_after": res2.RetryAfter.Seconds()})
		return
	}

	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	review.UserIP = ip
	review.CreatedAt = time.Now()

	if service.ContainsBannedWords(review.Content) || service.ContainsLinks(review.Content) {
		review.Status = "moderation"
	} else {
		review.Status = "approved"
	}

	service.SaveToHistory(ctx, review)
	c.JSON(http.StatusOK, gin.H{"status": review.Status})
}
