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

var Leagues *database.LEAGUE

func LeagueHandler(w http.ResponseWriter, r *http.Request) HTTPErrors.Error {
	urlPart := strings.Split(r.URL.Path, "/")
	fmt.Println("In League Handler")
	leagueID,_ := strconv.ParseInt(urlPart[4],10,64)


	switch r.Method {
	case "GET":
		if (urlPart[4] != "") {
			Leagues,_ = database.SelectLeague(leagueID)
			if len(urlPart) > 5 {
				switch urlPart[5] {
				case "results":
					fmt.Println("Results")
					groups,_ := database.SelectGroups("WHERE league_id=" + strconv.FormatInt(leagueID,10))
					var results []database.RESULT
					for i, group := range *groups {
						fmt.Println(i,group)
						tmp,_ := database.SelectResultLastByGroupId(group.ID)

						results = append(results, *tmp)
					}

					// Encode new structure to JSON format
					enc, err := json.Marshal(results)
					if err != nil {
						log.Fatalln(err)
					}

					// Gives JSON response for requests
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write(enc)

				case "groups":
					fmt.Println("Groups")
					groups,_ := database.SelectGroups("WHERE league_id=" + strconv.FormatInt(leagueID,10))
					// Encode new structure to JSON format
					enc, err := json.Marshal(groups)
					if err != nil {
						log.Fatalln(err)
					}

					// Gives JSON response for requests
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write(enc)
				default:
					return HTTPErrors.NewError("Bad Request", http.StatusBadRequest)
				}
			} else {
				// Encode new structure to JSON format
				enc, err := json.Marshal(Leagues)
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

