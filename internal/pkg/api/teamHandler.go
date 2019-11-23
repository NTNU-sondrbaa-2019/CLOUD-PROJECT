package api

import (
	"encoding/json"
	"fmt"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/HTTPErrors"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/view"
	"log"
	"net/http"
	"strings"
)

// Struct to store data gathered
type returnType struct {
	TeamName string
	Users []database.USER
}

var someValues []returnType

var user *database.USER
var team *[]database.GROUP
var grouping *[]database.GROUP_USER

func TeamHandler(w http.ResponseWriter, r *http.Request) HTTPErrors.Error {
	urlPart := strings.Split(r.URL.Path, "/")

	switch r.Method {
	case "GET":
		if (urlPart[4] != "") {
			// Search for team name urlPart[4]
			// This will now use ID, but in the future I would like to change this to something like by Name or Nickname
			team, _ = database.SelectGroups("WHERE name=\""+urlPart[4] + "\"")
			if len(urlPart)>5 {
				switch urlPart[5] {
				case "users":
					err := TeamUsersHandler(w,r)
					return err
				case "results":
					err := TeamResultsHandler(w,r)
					return err
				case "seasons":
					return HTTPErrors.NewError("Not Implemented", http.StatusNotImplemented)
				default:
					return HTTPErrors.NewError("Bad Request", http.StatusBadRequest)
				}
			} else {
				// Encode new structure to JSON format
				enc, err := json.Marshal(team)
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

func TeamUsersHandler(w http.ResponseWriter, r *http.Request) HTTPErrors.Error {
	fmt.Println("Finding Users in team...")
	// Since team is array this will return multiple teams
	for i, s := range *team {
		fmt.Println("Printing team...")
		fmt.Println(i,s)
		var tmpValues returnType
		grouping, _ = database.SelectGroupUserByGroupID(s.ID)
		tmpValues.TeamName = s.Name

		// Loop through usersgroups
		for n, ug := range *grouping {
			fmt.Println("Printing users")
			fmt.Println(n, ug)
			user, _ = database.SelectUserByID(ug.UserID)

			tmpValues.Users = append(tmpValues.Users, *user)
		}
		someValues = append(someValues, tmpValues)
	}

	// Encode new structure to JSON format
	enc, err := json.Marshal(someValues)
	if err != nil {
		log.Fatalln(err)
	}

	// Gives JSON response for requests
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(enc)

	return HTTPErrors.NewError("", 0)
}

func TeamResultsHandler(w http.ResponseWriter, r *http.Request) HTTPErrors.Error {

	return HTTPErrors.NewError("", 0)
}