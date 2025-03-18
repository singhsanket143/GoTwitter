package app

import (
	db "GoTwitter/db/repository"
	"GoTwitter/router"
	"log"
	"net/http"
	"time"
)

type Application struct {
	Config Config
	Store  db.Storage
}

type Config struct {
	Addr string
}

func (app *Application) Run() error {
	r := router.Mount() // no cycle here

	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      r,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server has started at %s", app.Config.Addr)

	return server.ListenAndServe()
}
