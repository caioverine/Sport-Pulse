package main

import (
	"log"
	"net/http"

	transporthttp "github.com/caioverine/sport-pulse/internal/transport/http"
)

func main() {
	router := transporthttp.NewRouter()

	log.Println("server listening on :8081")

	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatal(err)
	}
}
