package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sonalsguild/internal/router"

	"github.com/joho/godotenv"
)



type Config struct {
	Port string
}

type Application struct {
	Config Config
}

func (app *Application) Serve() error {
	 port := os.Getenv("PORT")
	 fmt.Println("API is listening on port",port);

	 server := &http.Server {
		Addr: fmt.Sprintf(":%s",port),
		Handler: router.Routes(),
	 }

	 return server.ListenAndServe()
}


func main() {
	fmt.Println("Starting...")
	// Load environment variables
	enverr := godotenv.Load(".env")
	if enverr != nil{
  		log.Fatalf("Error loading .env file: %s", enverr)
 	}

	


	// Server configutations
	cfg := Config {
		Port:os.Getenv("PORT"),
	}
	
	app := &Application {
		Config: cfg,
	}

	// Start server
	err := app.Serve()
	if err != nil {
		log.Fatal("Error starting server: ",err)
	}



    // DATABASE

	fmt.Println("Connected to the database, starting App")
}

