package handler

import (
	"net/http"
)

// HandleAPI - All requests to API endpoints will be handled here, follow testHandler for example
func HandleAPI(w http.ResponseWriter, r *http.Request, title string) {
	// Here we could make a html documentation file for our API for now it returns not found.
	http.NotFound(w, r)
}