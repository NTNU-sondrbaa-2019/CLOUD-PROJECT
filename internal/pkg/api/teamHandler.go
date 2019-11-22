package api

import (
	"encoding/json"
	"fmt"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/HTTPErrors"
	"log"
	"net/http"
	"strings"
)

type Team struct {
	Name 		string
	CreatedAt 	string
	Members 	[]User
}

func TeamHandler(w http.ResponseWriter, r *http.Request) HTTPErrors.Error {
	// Confirms to console that this handler was called
	fmt.Println("Team handler called.")
	urlPart := strings.Split(r.URL.Path, "/")

	switch r.Method {
	case "GET":
		// If it has more parameters than team/@
		if len(urlPart) > 5 {
			switch urlPart[5] {
			case "users":
				err := UsersTeamHandler(w,r)
				return err
			case "results":
				return HTTPErrors.NewError("Not Implemented", http.StatusNotImplemented)
			case "seasons":
				return HTTPErrors.NewError("Not Implemented", http.StatusNotImplemented)
			default:
				return HTTPErrors.NewError("Not Implemented", http.StatusNotImplemented)
			}
		} else {
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
		}
	default:
		return HTTPErrors.NewError("Invalid method", http.StatusBadRequest)
	}

	return HTTPErrors.NewError("", 0)
}

func UsersTeamHandler(w http.ResponseWriter, r *http.Request) HTTPErrors.Error {
	return HTTPErrors.NewError("Not Implemented", http.StatusNotImplemented)
}