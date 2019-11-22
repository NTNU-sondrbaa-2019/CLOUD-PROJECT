package handler

import (
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/view"
	"net/http"
	"strconv"
	"time"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/gauth"
)



type data struct {
	Title string // Title for page
	CurrentYear string // For the current year
	Username string // For username
	GoogleFetchData string
	GoogleClientID string
}

func HandleIndex(w http.ResponseWriter, r *http.Request, url string) {
	if url != "/" {
		http.NotFound(w, r)
	} else {
		logged := false // Doesnt check currently if actually logged in
		currentTime := time.Now()
		if !logged {
			// Page to load if logged in
			page := &data{Title: "Log in", CurrentYear: strconv.Itoa(currentTime.Year()), GoogleClientID: gauth.GoogleOauthConfig.ClientID}

			view.Render(w, "login", page)
		} else {
			// Page to load if not logged in
			page := &data{Title: "GR8ELO", CurrentYear: strconv.Itoa(currentTime.Year())}

			view.Render(w, "ucp", page)
		}
	}
}

