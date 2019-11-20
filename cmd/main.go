package main

import (
	"log"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^\\/(api\\/v1)\\/(\\S*)")

func main() {
	http.HandleFunc("/api/v1/", makeHandler(handlerTest))

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
