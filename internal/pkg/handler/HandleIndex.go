package handler

import (
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/gauth"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

var templates = template.Must(template.ParseFiles("web/static/login.html","web/static/ucp.html"))

type data struct {
	Title string // Title for page
	Date string // For the current year
	Username string // For username
	GoogleFetchData string
	GoogleClientID string
	GoogleRedirectURI string
}

func HandleIndex(w http.ResponseWriter, r *http.Request, url string) {
	if url != "/" {
		http.NotFound(w, r)
	} else {
		logged := false // Doesnt check currently if actually logged in
		currentTime := time.Now()

		// Gets the value od the sessionID cokkie from the users browser
		sessionID := gauth.GetCookieValueByName(r.Cookies(), "sessionID")
		// If the sessionID is not empty, we are logged in
		if sessionID != "" {
			logged = true
		}

		if !logged {
			// Page to load if not logged in
			page := &data{
				Title:             "Log in",
				Date:              strconv.Itoa(currentTime.Year()),
				GoogleFetchData:   gauth.GoogleOauthConfig.Scopes[0],
				GoogleClientID:    gauth.GoogleOauthConfig.ClientID,
				GoogleRedirectURI: gauth.GoogleOauthConfig.RedirectURL,
			}

			renderIndex(w, "login", page)
		} else {
			// Page to load if logged in
			page := &data{Title: "GR8ELO", Date: strconv.Itoa(currentTime.Year())}

			renderIndex(w, "ucp", page)
		}
	}
}

func renderIndex(writer http.ResponseWriter, s string, page interface{}) {
	// Assigns data to datapoints in html file
	err := templates.ExecuteTemplate(writer, s, page)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

