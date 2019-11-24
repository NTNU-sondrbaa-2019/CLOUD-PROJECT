package server

import (
	"log"
	"net/http"
	"os"
)

func Start() {
	port := os.Getenv("PORT")

	if port == "" {
		port = DEFAULT_PORT
	}

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
