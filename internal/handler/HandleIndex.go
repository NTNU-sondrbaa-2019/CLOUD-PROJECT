package handler

import (
	"html/template"
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

			RenderIndex(w, "login", page)
		} else {
			// Page to load if not logged in
			page := &data{Title: "GR8ELO", CurrentYear: strconv.Itoa(currentTime.Year())}

			RenderIndex(w, "ucp", page)
		}
	}
}

func RenderIndex(writer http.ResponseWriter, s string, page interface{}) {
	// Assigns data to datapoints in html file
	err := templates.ExecuteTemplate(writer, s, page)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

