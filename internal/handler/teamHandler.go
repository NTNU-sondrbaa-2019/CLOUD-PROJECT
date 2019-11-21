package handler

import (
	"encoding/json"
	"fmt"
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
		localTeam := Team{
			Name: urlPart[4],
			CreatedAt: "Date",
			Members: nil,
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
	default:
		// Not supported...
	}
}
