package main

import (
	"GoTwitter/app"
	config "GoTwitter/config/env"
	db "GoTwitter/db/repository"
	"fmt"
	"log"
)

func main() {

	fallback_port := ":3001"

	port := config.GetString("PORT", fallback_port)

	fmt.Println(port)

	cfg := app.Config{
		Addr: port,
	}

	store := db.NewMySQLStorage(nil)

	application := &app.Application{
		Config: cfg,
		Store:  store,
	}

	log.Printf("Server has started at %s", cfg.Addr)

	log.Fatal(application.Run())
}
