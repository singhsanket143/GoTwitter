package main

import (
	"GoTwitter/app"
	dbConfig "GoTwitter/config/db"
	config "GoTwitter/config/env"
	db "GoTwitter/db/repository"
	"fmt"
	"log"
)

func main() {

	// Load environment variables
	config.LoadEnv()

	fallback_port := ":3001"

	port := config.GetString("PORT", fallback_port)

	fmt.Println(port)

	cfg := app.Config{
		Addr: port,
		Db: dbConfig.DBConfig{
			Addr:               config.GetString("DB_ADDR", ""),
			MaxOpenConnections: config.GetInt("DB_MAX_OPEN_CONNECTIONS", 10),
			MaxIdleConnections: config.GetInt("DB_MAX_IDLE_CONNECTIONS", 10),
			MaxIdleTime:        config.GetInt("DB_MAX_IDLE_TIME", 10),
		},
	}

	dbConfig.SetupNewDbConn(cfg.Db.Addr, cfg.Db.MaxOpenConnections, cfg.Db.MaxIdleConnections, cfg.Db.MaxIdleTime)

	store := db.NewMySQLStorage(nil)

	application := &app.Application{
		Config: cfg,
		Store:  store,
	}

	log.Printf("Server has started at %s", cfg.Addr)

	log.Fatal(application.Run())
}
