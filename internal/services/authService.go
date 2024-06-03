package services

import (
	"context"
	"database/sql"
	"log"
)


func CheckUserExists(email string) bool {
    // Assuming db is your *sql.DB connection pool
    var count int
    query := `SELECT email FROM users WHERE email = $1`
    err := db.QueryRow(query, email).Scan(&count)
    if err!= nil && err!= sql.ErrNoRows {
        log.Fatalf("Failed to query user existence: %v", err)
    }
    return count > 0
}

func (s *User) RegisterUser(user User) (*User,error) {
	ctx, cancel := context.WithTimeout(context.Background(),dbTimeout)

	defer cancel()
	query := `
		INSERT INTO users (display_name, email) 
		VALUES ($1, $2, $3) returning *
	`
	 
	_, err := db.ExecContext(
		ctx,
		query,
		user.DisplayName,
		user.Email)
	
	if err != nil {
		return nil, err
	}
	

	return &user,nil
	
}