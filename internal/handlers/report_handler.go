package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maksroxx/ReviewGuard/internal/db"
)

func GetSpamReviewsHandler(rep db.PostgresRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		reviews, err := rep.GetReviewsByStatus(ctx, "moderation")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve moderation reviews"})
			return
		}
		c.JSON(http.StatusOK, reviews)
	}
}

func GetReviewsByIPHandler(rep db.PostgresRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		ip := c.Query("ip")
		if ip == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "missing ip parameter"})
			return
		}
		reviews, err := rep.GetReviewsByIP(ctx, ip)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve reviews by IP"})
			return
		}
		c.JSON(http.StatusOK, reviews)
	}
}
