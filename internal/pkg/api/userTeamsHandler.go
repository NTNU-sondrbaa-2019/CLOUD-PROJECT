package api

import (
	"encoding/json"
	"fmt"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/HTTPErrors"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
	"log"
	"net/http"
)

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
