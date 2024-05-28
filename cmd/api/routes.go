package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() *chi.Mux {
	// Initialize a new router
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	// Routes
	router.Get("/v1/healthcheck", app.healthcheckHandler)
	router.Post("/v1/apartments", app.createApartmentHandler)
	router.Get("/v1/apartments/{id}", app.getApartmentHandler)
	return router
}
