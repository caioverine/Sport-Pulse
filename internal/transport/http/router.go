package http

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /events", eventsHandler)
	mux.HandleFunc("GET /events/stream", eventsStreamHandler)

	return mux
}
