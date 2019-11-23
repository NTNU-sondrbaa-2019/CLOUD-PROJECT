package api

import (
    "encoding/json"
	"fmt"
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
			userID,_ := strconv.ParseInt(urlPart[4], 10, 64)
			user, _ = database.SelectUserByID(userID)
			if len(urlPart) > 5 {
				switch urlPart[5] {
				case "results":
					return HTTPErrors.NewError("Not Implemented", http.StatusNotImplemented)
				case "teams":
					err := UserTeamsHandler(w,r)
					return err
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

type userTeams struct {
	Username string
	UserID int64
	Groups []database.GROUP
}

func UserTeamsHandler(w http.ResponseWriter, r *http.Request) HTTPErrors.Error {
	fmt.Println("Finding Users in group...")
	// Since group is array this will return multiple teams
	group_user, _ = database.SelectGroupUserByUserID(user.ID)

	var ut userTeams

	ut.UserID = user.ID
	ut.Username = user.Name


	for n, ug := range *group_user {
		fmt.Println("Printing users")
		fmt.Println(n, ug)

		group, _ = database.SelectGroup(ug.GroupID)
		ut.Groups = append(ut.Groups, *group)
	}

	// Encode new structure to JSON format
	enc, err := json.Marshal(ut)
	if err != nil {
		log.Fatalln(err)
	}

	// Gives JSON response for requests
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(enc)

	return HTTPErrors.NewError("", 0)
}