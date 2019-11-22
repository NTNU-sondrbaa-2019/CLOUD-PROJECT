package user

import (
    "encoding/json"
    "github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/view"
    "log"
    "net/http"
    "strings"
)

type User struct {
	Name string
	Email string
	Registered string
	LastOnline string
	ID int
}

func UserHandler(w http.ResponseWriter, r *http.Request, title string) {

	urlPart := strings.Split(r.URL.Path, "/")
	switch r.Method {
	case "GET":
		// If it has more parameters than team/@
		if len(urlPart) > 5 {
			switch urlPart[5] {
			case "results":
				view.ErrorPage(w, "Not implemented", http.StatusNotImplemented)
			case "teams":
				view.ErrorPage(w, "Not implemented", http.StatusNotImplemented)
			default:
				view.ErrorPage(w, "Not Found", http.StatusNotFound)
			}
		} else {
			// Search for team name urlPart[4]
			user := User{
				Name: "Name",
			}

			// Encode new structure to JSON format
			enc, err := json.Marshal(user)
			if err != nil {
				log.Fatalln(err)
			}

			// Gives JSON response for requests
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(enc)
		}
	default:
		view.ErrorPage(w, "Not Implemented", http.StatusNotImplemented)
	}
}
