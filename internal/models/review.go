package models

import "time"

type Review struct {
	ID        string    `json:"id"`
	UserIP    string    `json:"user_ip"`
	Content   string    `json:"content"`
	Status    string    `json:"status"` // approved | moderation
	CreatedAt time.Time `json:"created_at"`
}
