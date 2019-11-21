package handler

import (
	"fmt"
	"net/http"
)

// HandleAPI - All requests to API endpoints will be handled here, follow testHandler for example
func HandleAPI(w http.ResponseWriter, r *http.Request, title string) {
		http.HandleFunc("/api/v1/test/", MakeHandler(testHandler)) // Should be removed in final version

		// Here we could make a html documentation file for our API for now it returns not found.
		http.NotFound(w, r)
}

func testHandler(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Println("Test Handler was called")
}