package index

import (
	"fmt"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/HTTPErrors"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/gauth"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/view"
	"net/http"
	"strconv"
	"time"
)

type data struct {
	Title       string // Title for page
	CurrentYear string // For the current year
	Username    string // For username
}

func HandleIndex(w http.ResponseWriter, r *http.Request) HTTPErrors.Error {
	url := r.URL.Path
	fmt.Println(url)
	if url == "/" {
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
				Title:       "Login - gr8elo.com",
				CurrentYear: strconv.Itoa(currentTime.Year()),
				Username:    "Unknown Username",
			}

			view.Render(w, "login", page)
		} else {
			// Page to load if logged in
			page := &data{Title: "Homepage - gr8elo.com", CurrentYear: strconv.Itoa(currentTime.Year())}

			view.Render(w, "ucp", page)
		}
	} else {
		view.ErrorPage(w, "Not found", http.StatusNotFound)
	}
	return HTTPErrors.NewError("", 0)
}
