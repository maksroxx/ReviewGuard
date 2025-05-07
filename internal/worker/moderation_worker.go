package worker

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/maksroxx/ReviewGuard/internal/db"
	"github.com/maksroxx/ReviewGuard/internal/models"
	"github.com/maksroxx/ReviewGuard/internal/service"
	"github.com/redis/go-redis/v9"
)

func StartModerationWorker(redisClient *redis.Client, repo *db.PostgresRepository) {
	ctx := context.Background()

	err := redisClient.XGroupCreateMkStream(ctx, "moderation_stream", "moderators", "$").Err()
	if err != nil && err.Error() != "BUSYGROUP Consumer Group name already exists" {
		log.Fatalf("Failed to create consumer group: %v", err)
	}

	for {
		streams, err := redisClient.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    "moderators",
			Consumer: "worker-1",
			Streams:  []string{"moderation_stream", ">"},
			Count:    10,
			Block:    5 * time.Second,
		}).Result()

		if err != nil && err != redis.Nil {
			log.Printf("Error reading from stream: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		for _, stream := range streams {
			for _, message := range stream.Messages {
				raw := message.Values["data"].(string)

				var review models.Review
				if err := json.Unmarshal([]byte(raw), &review); err != nil {
					log.Printf("Failed to parse review: %v", err)
					continue
				}

				if service.ContainsBannedWords(review.Content) || service.ContainsLinks(review.Content) {
					review.Status = "moderation"
				} else {
					review.Status = "approved"
				}

				if err := repo.UpdateReviewStatus(ctx, review.ID, review.Status); err != nil {
					log.Printf("Failed to update review status: %v", err)
					continue
				}

				redisClient.XAck(ctx, "moderation_stream", "moderators", message.ID)
			}
		}
	}
}
