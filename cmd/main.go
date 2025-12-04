package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/qlfzn/goffee-club/cmd/api"
	"github.com/qlfzn/goffee-club/internal/repository"
)

func main() {
	logger := log.New(os.Stdout, "[GOFFEE-CLUB] ", log.LstdFlags|log.Lshortfile)

	err := godotenv.Load()
	if err != nil {
		logger.Fatalf("error loading .env file: %s", err)
	}

	// db config
	dbConf := repository.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	cfg := api.Config{
		Addr: ":8080",
		DB: &dbConf,
	}

	// connect to db
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := repository.NewDBPool(ctx, dbConf)
	if err != nil {
		logger.Fatalf("error initiating database connection %s", err)
	}
	defer repository.Close(db)
	logger.Print("successfully connected to database")

	app := api.Application{
		Config: cfg,
		DB:     db,
		Logger: logger,
	}

	r := app.Mount()

	err = app.Run(r)
	if err != nil {
		logger.Fatalf("error running server: %s", err)
	}
}
