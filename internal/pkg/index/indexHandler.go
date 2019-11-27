package index

import (
	"fmt"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/HTTPErrors"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
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

		signed_in := false // Doesnt check currently if actually signed_in in
		currentTime := time.Now()

		// Gets the value od the sessionID cokkie from the users browser
		sessionID := gauth.GetCookieValueByName(r.Cookies(), "sessionID")
		// If the sessionID is not empty, we are signed_in in
		if sessionID != "" {
			signed_in = true
		}

		if !signed_in {

			// Page to load if not signed_in in
			page := &data{
				Title:       "Login - gr8elo.com",
				CurrentYear: strconv.Itoa(currentTime.Year()),
				Username:    "Guest",
			}

			view.Render(w, "login", page)

		} else {

			var user *database.USER

			user_id, err := gauth.GetUserIDFromSessionID(sessionID)

			if err != nil {

				return HTTPErrors.NewError("Couldn't fetch user id from session", http.StatusInternalServerError)

			} else {

				user, err = database.SelectUserByID(user_id)

				if err != nil {

					return HTTPErrors.NewError("Couldn't fetch user with user id (from session)", http.StatusInternalServerError)

				}

			}

			// Page to load if signed_in in
			page := &data{
				Title:       "Homepage - gr8elo.com",
				CurrentYear: strconv.Itoa(currentTime.Year()),
				Username:    user.Name,
			}

			view.Render(w, "ucp", page)

		}

	} else {

		return HTTPErrors.NewError("Not found", http.StatusNotFound)

	}

	return HTTPErrors.NewError("Something went wrong", http.StatusInternalServerError)

}
