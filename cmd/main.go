package main

import (
	"log"
	"net/http"

	"github.com/maksroxx/ReviewGuard/internal/api"
)

func main() {
	http.HandleFunc("/review", api.ReviewHandler)
	log.Println("ReviewGuard :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
