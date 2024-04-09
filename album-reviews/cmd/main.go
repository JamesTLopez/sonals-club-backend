package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sonalsguild/internal/db"

	"github.com/joho/godotenv"
)



type Config struct {
	Port string
}

type Application struct {
	Config Config
	router http.Handler

}

func (app *Application) Serve() error {
	 port := os.Getenv("PORT")
	 fmt.Println("API is listening on port",port);

	 server := &http.Server {
		Addr: fmt.Sprintf(":%s",port),
		Handler: app.router,
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



	connectionString := os.Getenv("DATABASE_URL")

	dbConn , databaseErr := db.ConnectPostgres(connectionString)

	if databaseErr != nil {
		log.Fatal("Cannot connect to database")
	}

	defer dbConn.DB.Close()

	// Server configutations
	cfg := Config {
		Port:os.Getenv("PORT"),
	}
	
	app := &Application {
		Config:cfg,
		// TODO: models
	}

	// Start server
	err := app.Serve()
	if err != nil {
		log.Fatal("Error starting server: ",err)
	}



    // DATABASE

	fmt.Println("Connected to the database, starting App")
}

