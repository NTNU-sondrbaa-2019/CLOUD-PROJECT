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

func GroupHandler(w http.ResponseWriter, r *http.Request) HTTPErrors.Error {
	urlPart := strings.Split(r.URL.Path, "/")
	switch r.Method {
	case "GET":
		if (urlPart[4] != "") {
			// Search for group name urlPart[4]
			// This will now use ID, but in the future I would like to change this to something like by Name or Nickname

			// SQL INJECTION - Disabled by Sondre
			// groups, _ = database.SelectGroups("WHERE name=\""+urlPart[4] + "\"")

			var groups []database.GROUP

			fmt.Println(groups)
			if len(urlPart)>5 {
				switch urlPart[5] {
				case "users":
					err := GroupUsersHandler(w,r)
					return err
				case "results":
					err := GroupResultsHandler(w,r)
					return err
				case "seasons":
					return HTTPErrors.NewError("Not Implemented", http.StatusNotImplemented)
				default:
					return HTTPErrors.NewError("Bad Request", http.StatusBadRequest)
				}
			} else {
				// Encode new structure to JSON format
				enc, err := json.Marshal(groups)
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

func GroupResultsHandler(w http.ResponseWriter, r *http.Request) HTTPErrors.Error {
	fmt.Println("Finding group results...")

	var tmpGroupResults groupRes
	var groupResults []groupRes

	for i, g := range *groups {
		fmt.Println(i,g)

		tmpGroupResults.GroupName = g.Name
		tmpGroupResults.Results, _  = database.SelectResults("WHERE group_id=" + strconv.FormatInt(g.ID, 10))

		groupResults = append(groupResults, tmpGroupResults)
	}

	// Encode new structure to JSON format
	enc, err := json.Marshal(groupResults)
	if err != nil {
		log.Fatalln(err)
	}

	// Gives JSON response for requests
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(enc)

	return HTTPErrors.NewError("", 0)
}