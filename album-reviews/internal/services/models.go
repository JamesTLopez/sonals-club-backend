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
	password string 
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


type Song struct {
	ID string `json:"id"`
	userID string
	Name string `json:"name"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}