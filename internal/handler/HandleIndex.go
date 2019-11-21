package handler

import (
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request, url string) {
	if url != "/" {
		http.NotFound(w, r)
	} else {
		// Check if logged in
		// if logged in
			// Load user page
		// else if load login page
	}
}

