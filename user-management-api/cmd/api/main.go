package main

import (
	"user-management-api/internal/app"
	"user-management-api/internal/config"
)

func main() {

	// init config
	cfg := config.NewConfig()

	// init application
	application := app.NewApplication(cfg)

	//start server
	if err := application.Run(); err != nil {
		panic("error start server")
	}
}
