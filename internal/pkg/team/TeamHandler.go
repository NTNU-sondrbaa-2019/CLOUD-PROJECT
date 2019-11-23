package team

import (
	"encoding/json"
	"fmt"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/view"
	"log"
	"net/http"
	"strings"
)

type Team struct {
	Name 		string
	CreatedAt 	string
	Members 	[]User
}

type User struct {
	Username string
}

func TeamHandler(w http.ResponseWriter, r *http.Request, title string) {
	// Confirms to console that this handler was called
	fmt.Println("Team handler called.")
	urlPart := strings.Split(r.URL.Path, "/")

	switch r.Method {
	case "GET":
		// If it has more parameters than team/@
		if len(urlPart) > 5 {
			switch urlPart[5] {
			case "users":
				UsersTeamHandler(w,r,title)
			case "results":
				view.ErrorPage(w, "Not implemented", http.StatusNotImplemented)
			case "seasons":
				view.ErrorPage(w, "Not implemented", http.StatusNotImplemented)
			default:
				view.ErrorPage(w, "Not Found", http.StatusNotFound)
			}
		} else if (len(urlPart) == 5) {
			// Search for team name urlPart[4]
			localTeam := Team{
				Name: "Name",
			}

			// Encode new structure to JSON format
			enc, err := json.Marshal(localTeam)
			if err != nil {
				log.Fatalln(err)
			}

			// Gives JSON response for requests
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(enc)
		} else {
			view.ErrorPage(w, "Not implemented", http.StatusNotImplemented)
		}
	default:
		view.ErrorPage(w, "Not implemented", http.StatusNotImplemented)
	}
}

func UsersTeamHandler(w http.ResponseWriter, r *http.Request, title string) {
	view.ErrorPage(w, "Not implemented", http.StatusNotImplemented)
}