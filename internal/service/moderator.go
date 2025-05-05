package service

import (
	"fmt"
	"sync"

	"github.com/maksroxx/ReviewGuard/internal/filter"
	"github.com/maksroxx/ReviewGuard/internal/models"
	"github.com/maksroxx/ReviewGuard/internal/spam"
)

var previousReviews = make(map[string]bool)
var mu sync.Mutex

func Moderate(review *models.Review) {
	mu.Lock()
	defer mu.Unlock()

	switch {
	case filter.ContainsBannedWords(review.Content),
		filter.ContainsLinks(review.Content),
		filter.IsDuplicate(review.Content, previousReviews),
		spam.IsSuspicious(review.UserIP):

		review.Status = "moderation"
	default:
		review.Status = "approved"
		previousReviews[review.Content] = true
	}
	fmt.Printf("Review: %+v\n", review)
}
