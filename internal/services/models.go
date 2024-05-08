package services

import "time"

type JsonResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitresponse"`
}


type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	password string `json:"password"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


type Song struct {
	ID string `json:"id"`
	// userID string `json:"user_id"` 
	Name string `json:"name"`
	Labels string `json:"labels"`
	Description string `json:"description"`
	Duration int `json:"duration"`
	Color string `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Sample struct {
	ID string `json:"id"`
	userId string `json:"user_id"`
	songId string `json:"song_id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}