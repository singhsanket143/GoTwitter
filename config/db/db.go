package config

import (
	config "GoTwitter/config/env"
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func SetupNewDbConn() (*sql.DB, error) {

	addr := config.GetString("DB_ADDR", "")
	maxOpenConnections := config.GetInt("DB_MAX_OPEN_CONNECTIONS", 10)
	maxIdleConnections := config.GetInt("DB_MAX_IDLE_CONNECTIONS", 10)
	maxIdleTime := config.GetInt("DB_MAX_IDLE_TIME", 10)

	log.Printf("Connecting to database at %s", addr)

	db, err := sql.Open("mysql", addr)

	if err != nil {
		log.Fatalf("Error opening database: %v", err)
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConnections)
	db.SetMaxIdleConns(maxIdleConnections)
	db.SetConnMaxIdleTime(time.Minute * time.Duration(maxIdleTime))

	err = db.Ping()
	log.Printf("Trying to ping database at %s", addr)
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Verify the connection
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	log.Printf("Connected to database at %s", addr)

	return db, nil
}
