package handler

import (
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
}

func HandleIndex(w http.ResponseWriter, r *http.Request, url string) {
	if url != "/" {
		http.NotFound(w, r)
	} else {
		logged := false // Doesnt check currently if actually logged in
		currentTime := time.Now()
		if !logged {
			// Page to load if logged in
			page := &data{Title: "Log in", Date: strconv.Itoa(currentTime.Year())}

			renderIndex(w, "login", page)
		} else {
			// Page to load if not logged in
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

