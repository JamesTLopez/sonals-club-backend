package application

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App {
		router: loadRoutes(),
	}
	return app
}

func ( a *App ) Start(ctx context.Context) error {
	conn, dberr := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if dberr != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", dberr)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	
	server := &http.Server {
		Addr: ":8080",
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil

}