package handler

import (
	"fmt"
	"net/http"
)

type Team struct {
	Name 		string
	CreatedAt 	string
	members 	[]User
}

type User struct {

}

func TeamHandler(w http.ResponseWriter, r *http.Request, title string) {
	// Confirms to console that this handler was called
	fmt.Println("Team handler called.")

	switch r.Method {
	case "GET":
		
	default:
		// Not supported...
	}
}
