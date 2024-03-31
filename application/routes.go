package application

import (
	"net/http"

	"sonalsguild/api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger);

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK);
	})
	
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK);
	})

	router.Route("/review",loadReviewRoutes)
	return router
}


func loadReviewRoutes ( router chi.Router ) {
	reviewHandler := &handlers.Order{}

	router.Post("/",reviewHandler.CreateReview)
	router.Get("/",reviewHandler.ReviewList)
	router.Get("/{id}",reviewHandler.GetReviewById)
	router.Put("/{id}",reviewHandler.UpdateReviewById)
	router.Delete("/{id}",reviewHandler.DeleteReviewById)
}