package app

import (
    "net/http"
    "time"
    "log"
    "GoTwitter/router"
)

type Application struct {
    Config Config
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
