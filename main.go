package main

import (
	"context"
	"fmt"
	"log"
	"sonalsguild/api/application"
)

func main() {
	log.Println("Starting...")
	app := application.New()

	err := app.Start(context.TODO())

	if err != nil {
		fmt.Printf("Error starting server %v",err)
	}
}

