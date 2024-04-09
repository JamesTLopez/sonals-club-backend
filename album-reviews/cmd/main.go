package main

import (
	"context"
	"fmt"
	"log"
	"os"
	application "sonalsguild/internal/api"

	"github.com/jackc/pgx/v5"
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

	conn, errs := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if errs != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", errs)
		os.Exit(1)
	}
	defer conn.Close(context.Background())


	fmt.Println("Connected to the database, starting App")
	
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Printf("Error starting server %v",err)
	}
}

