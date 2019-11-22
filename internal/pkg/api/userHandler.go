package api

import (
    "encoding/json"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/HTTPErrors"
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

func UserHandler(w http.ResponseWriter, r *http.Request) HTTPErrors.Error {

	urlPart := strings.Split(r.URL.Path, "/")
	switch r.Method {
	case "GET":
		// If it has more parameters than team/@
		if len(urlPart) > 5 {
			switch urlPart[5] {
			case "results":
				return HTTPErrors.NewError("Not Implemented", http.StatusNotImplemented)
			case "teams":
				return HTTPErrors.NewError("Not Implemented", http.StatusNotImplemented)
			default:
				return HTTPErrors.NewError("Not Implemented", http.StatusNotImplemented)
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
		return HTTPErrors.NewError("Invalid method", http.StatusBadRequest)
	}

	return HTTPErrors.NewError("", 0)
}
