package models

import "time"

type Song struct {
	ID string `json:"id"`
	// userID string `json:"user_id"` 
	Name *string `json:"song_name"`
	Labels string `json:"labels"`
	Description *string `json:"description"`
	Duration int `json:"duration"`
	Color string `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
