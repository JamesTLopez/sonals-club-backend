package services

import (
	"context"
	"time"
)

type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func (s *User) RegisterUser(user User) (*User,error) {
	ctx, cancel := context.WithTimeout(context.Background(),dbTimeout)

	defer cancel()
	query := `
		INSERT INTO users (username, password, email) 
		VALUES ($1, $2, $3) returning *
	`
	 
	_, err := db.ExecContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email)
	
	if err != nil {
		return nil, err
	}
	

	return &user,nil

}