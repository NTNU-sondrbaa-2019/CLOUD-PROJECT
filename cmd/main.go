package main

import (
	"log"
	"net/http"
)

func main() {
	// 2 Options, either all handle requests are made here with new handlers
	// Or each endpoint can be handled via switch in a "mainHandler"
	//http.HandleFunc("/api/v1/", internal.MakeHandler(someHandler))

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
