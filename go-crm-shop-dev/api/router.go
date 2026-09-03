package api

import (
	"net/http"

	"go-crm-shop-dev/api/handler"
	"go-crm-shop-dev/api/middleware"
)

// NewRouter builds the HTTP API router.
func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handler.Health)

	return middleware.RequestID(mux)
}
