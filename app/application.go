package app

import (
	dbConfig "GoTwitter/config/db"
	db "GoTwitter/db/repository"
	"GoTwitter/router"
	"net/http"
	"time"
)

type Application struct {
	Config Config
	Store  db.Storage
}

type Config struct {
	Addr string
	Db   dbConfig.DBConfig
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

	return server.ListenAndServe()
}
