package services

import (
	"context"
	"fmt"
	"log"
)


func (a *Authorization) CheckUserExists(email string) bool {
	ctx, cancel := context.WithTimeout(context.Background(),dbTimeout)
	defer cancel()
    
	sql, _, err := psql.Select("email").From("users").Where("email IN ($1)").ToSql()
	if err != nil {
		log.Fatalf("Failed to generate sql from squireel: %v",err)
	}


		
	var user User
	row :=  db.QueryRowContext(ctx,sql,email)

	if row == nil {
		fmt.Println("User does not exist")
		return false
	}


	err = row.Scan(&user.Email)
	if err!= nil {
		fmt.Println("User does not exist:", err)
		return false
	}

    return true
}

func (a *Authorization) RegisterUser(spotify_id string, display_name string, email string) (bool,error) {
	ctx, cancel := context.WithTimeout(context.Background(),dbTimeout)

	defer cancel()
	query := `
		INSERT INTO users ( display_name, email) 
		VALUES ($1, $2) returning *
	`
	 
	_, err := db.ExecContext(
		ctx,
		query,
		display_name,
		email)
	
	if err != nil {
		return false, err
	}
	

	return true,nil
	
}