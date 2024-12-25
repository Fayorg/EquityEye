package server

import "net/http"

func NewServer() http.Handler {
	mux := http.NewServeMux()

	var handler http.Handler = mux
	// handler = middlewares.LoggingMiddleware(handler)

	addRoutes(mux)

	return handler
}
