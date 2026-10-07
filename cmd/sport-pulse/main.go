package main

import (
	"log"
	"net/http"

	transporthttp "github.com/caioverine/sport-pulse/internal/transport/http"
)

func main() {
	router := transporthttp.NewRouter()

	log.Println("server listening on :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
