package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maksroxx/ReviewGuard/internal/db"
	"github.com/maksroxx/ReviewGuard/internal/service"
)

func ReportStatsHandler(svc service.CacheSerice, rep db.PostgresRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		period := c.DefaultQuery("period", "hour")

		stats, err := svc.GetCachedIPStats(ctx, period)
		if err != nil {
			stats, err = rep.GetIPStats(ctx, period)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve stats"})
				return
			}
			_ = svc.CacheIPStats(ctx, period, stats)
		}
		c.JSON(http.StatusOK, gin.H{
			"period": period,
			"stats":  stats,
		})
	}
}
