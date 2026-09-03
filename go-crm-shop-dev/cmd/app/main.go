package main

import (
	"log"
	"net/http"
	"os"

	"go-crm-shop-dev/api"
)

func main() {
	address := os.Getenv("APP_ADDR")
	if address == "" {
		address = ":8080"
	}

	log.Printf("CRM shop API listening on %s", address)
	if err := http.ListenAndServe(address, api.NewRouter()); err != nil {
		log.Fatal(err)
	}
}
