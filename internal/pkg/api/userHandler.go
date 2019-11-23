package api

import (
    "encoding/json"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/HTTPErrors"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/view"
	"log"
    "net/http"
	"strconv"
	"strings"
)

func UserHandler(w http.ResponseWriter, r *http.Request) HTTPErrors.Error {
	urlPart := strings.Split(r.URL.Path, "/")

	switch r.Method {
	case "GET":
		if (urlPart[4] != "") {
			// Search for team name urlPart[4]
			// This will now use ID, but in the future I would like to change this to something like by Name or Nickname
			userID,_ := strconv.ParseInt(urlPart[4], 10, 64)
			user, _ = database.SelectUserByID(userID)
			if len(urlPart) > 5 {
				switch urlPart[5] {
				case "users":
					err := TeamUsersHandler(w, r)
					return err
				case "results":
					err := TeamResultsHandler(w, r)
					return err
				case "seasons":
					return HTTPErrors.NewError("Not Implemented", http.StatusNotImplemented)
				default:
					return HTTPErrors.NewError("Bad Request", http.StatusBadRequest)
				}
			} else {
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
		} else {
			return HTTPErrors.NewError("Bad Request", http.StatusBadRequest)
		}
	default:
		view.ErrorPage(w, "Not implemented", http.StatusNotImplemented)
	}

	return HTTPErrors.NewError("", 0)
}