package router

import (
	"net/http"
	"sonalsguild/internal/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Routes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options {
        AllowedOrigins: []string{"http://*", "https://*"}, // TODO: for security, change such that it targets published origins
        AllowedMethods: []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders: []string{"Link"},
        AllowCredentials: true,
        MaxAge: 300,
	}) )

	// Songs Routes
	router.Get("/api/v1/songs",controllers.GetAllSongs)
	router.Get("/api/v1/songs/{id}",controllers.GetSongById)
	router.Post("/api/v1/songs",controllers.CreateSong)
	router.Put("/api/v1/songs/{id}",controllers.UpdateSong)
	router.Delete("/api/v1/songs/{id}",controllers.DeleteSong)


	// Samples Routes
	router.Get("/api/v1/samples", controllers.GetAllSamples);
	router.Post("/api/v1/samples",controllers.CreateSample);
	
	return router
}