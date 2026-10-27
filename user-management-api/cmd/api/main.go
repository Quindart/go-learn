package main

import (
	"log"
	"user-management-api/internal/app"
	"user-management-api/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found; using environment variables")
	}

	// init config
	cfg := config.NewConfig()

	// init application
	application := app.NewApplication(cfg)

	//start server
	if err := application.Run(); err != nil {
		panic("error start server")
	}
}
