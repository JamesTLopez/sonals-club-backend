package main

import (
	"context"
	"fmt"
	"log"
	application "sonalsguild/internal/api"
	database "sonalsguild/internal/db"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting...")
	app := application.New()
	// Load environment variables
	enverr := godotenv.Load(".env")
	if enverr != nil{
  		log.Fatalf("Error loading .env file: %s", enverr)
 	}

	connPool,errorDbConnection := pgxpool.NewWithConfig(context.Background(), database.Config())
	if errorDbConnection != nil {
		log.Fatal("Error while creating connection to the database!!")
	} 

	connection, errAquire := connPool.Acquire(context.Background())
	if errAquire != nil {
		log.Fatal("Error while acquiring connection from the database pool!!")
	} 
	defer connection.Release()

	errPing := connection.Ping(context.Background())
	if errPing != nil{
		log.Fatal("Could not ping database")
	}

	fmt.Println("Connected to the database, starting App")
	
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Printf("Error starting server %v",err)
	}
}

