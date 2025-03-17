package main

import (
	"GoTwitter/app"
	config "GoTwitter/config/env"
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

	application := &app.Application{
		Config: cfg,
	}

	log.Printf("Server has started at %s", cfg.Addr)

	log.Fatal(application.Run())
}
