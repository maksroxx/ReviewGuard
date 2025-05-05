package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/maksroxx/ReviewGuard/internal/models"
	"github.com/maksroxx/ReviewGuard/internal/service"
)

func ReviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var review models.Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	review.CreatedAt = time.Now()

	service.Moderate(&review)

	response, _ := json.Marshal(map[string]string{
		"status": review.Status,
	})
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
