package services

import (
	"context"
)




func (s *User) RegisterUser(user User) (*User,error) {
	ctx, cancel := context.WithTimeout(context.Background(),dbTimeout)

	defer cancel()
	query := `
		INSERT INTO users (display_name, password, email) 
		VALUES ($1, $2, $3) returning *
	`
	 
	_, err := db.ExecContext(
		ctx,
		query,
		user.DisplayName,
		user.password,
		user.Email)
	
	if err != nil {
		return nil, err
	}
	

	return &user,nil

}