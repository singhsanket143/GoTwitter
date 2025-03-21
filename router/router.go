package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Routes interface {
	Register(r chi.Router)
}

func Mount(tweetRouter Routes) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)

	tweetRouter.Register(router)

	// router.Route("/v1", func(router chi.Router) {
	// 	router.Get("/ping", handlers.PingHandler)
	// })

	return router
}
