package router

import (
	"net/http"
	"sonalsguild/internal/controllers"

	sMiddle "sonalsguild/internal/middleware"

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

	router.Group(func(r chi.Router) {
		r.Use(sMiddle.VerifyToken)
		// Songs Routes
		r.Get("/api/v1/songs",controllers.GetAllSongs)
		r.Get("/api/v1/songs/{id}",controllers.GetSongById)
		r.Post("/api/v1/songs",controllers.CreateSong)
		r.Put("/api/v1/songs/{id}",controllers.UpdateSong)
		r.Delete("/api/v1/songs/{id}",controllers.DeleteSong)


		// Samples Routes
		r.Get("/api/v1/samples", controllers.GetAllSamples);
		r.Post("/api/v1/samples",controllers.CreateSample);
		r.Put("/api/v1/samples/{id}", controllers.UpdateSample)
		r.Delete("/api/v1/samples/{id}", controllers.DeleteSample);
	})
	


	// AuthRouters Spotify
	router.Get("/auth/{provider}",controllers.GetAutheniticateSpotify);
	router.Get("/auth/{provider}/callback",controllers.GetAuthCallbackSpotify);

	
	return router
}