package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Team struct {
	Name 		string
	CreatedAt 	string
	Members 	[]User
}

type User struct {
	Username string
}

type Page struct{

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
				ErrorPage(w, "Not implemented", http.StatusNotImplemented)
			case "seasons":
				ErrorPage(w, "Not implemented", http.StatusNotImplemented)
			default:
				ErrorPage(w, "Not Found", http.StatusNotFound)
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
		// Not supported...
	}
}

func UsersTeamHandler(w http.ResponseWriter, r *http.Request, title string) {
	ErrorPage(w, "Not implemented", http.StatusNotImplemented)
}

func ErrorPage(w http.ResponseWriter,errorMsg string, code int) {
	type err struct {
		ErrorMsg string
		ErrorCode int
		CurrentYear string
	}
	error := &err{ErrorMsg: errorMsg, ErrorCode: code, CurrentYear: strconv.Itoa(time.Now().Year())}

	RenderIndex(w, "error", error)
}