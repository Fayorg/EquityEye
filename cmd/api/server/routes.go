package server

import (
	"EquityEye/internal/http/handlers"
	"net/http"
)

func addRoutes(mux *http.ServeMux) {
	mux.Handle("GET /ping", handlers.PingHandler())

	// Catch-all route
	mux.Handle("/", http.NotFoundHandler())
}
