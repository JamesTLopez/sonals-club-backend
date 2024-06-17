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
	})
	


	// AuthRouters Spotify
	router.Get("/auth/{provider}",controllers.GetAutheniticateSpotify);
	router.Get("/auth/{provider}/callback",controllers.GetAuthCallbackSpotify);

	
	return router
}