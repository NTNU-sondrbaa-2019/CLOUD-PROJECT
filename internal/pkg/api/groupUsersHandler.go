package api

import (
	"encoding/json"
	"fmt"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/HTTPErrors"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
	"log"
	"net/http"
)

func GroupUsersHandler(w http.ResponseWriter, r *http.Request) HTTPErrors.Error {
	fmt.Println("Finding Users in group...")
	// Since groups is array this will return multiple teams
	for i, s := range *groups {
		fmt.Println("Printing group...")
		fmt.Println(i,s)
		var tmpValues teamsUser
		group_user, _ = database.SelectGroupUserByGroupID(s.ID)
		tmpValues.TeamName = s.Name

		fmt.Println("Group name", s.Name)

		// Loop through usersgroups
		for n, ug := range *group_user {
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
