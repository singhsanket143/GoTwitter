package router

import (
    "github.com/go-chi/chi/v5"
    "GoTwitter/handlers"
	"github.com/go-chi/chi/v5/middleware"
)

func Mount() *chi.Mux {
    router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)

	router.Route("/v1", func (router chi.Router) {
		router.Get("/ping", handlers.PingHandler)
	})

    return router
}
