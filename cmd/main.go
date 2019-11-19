package main

import (
	"fmt"
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal"
	"net/http"

)

func getMemberPrint(){

}



func main() {

	type Test struct {
		Name string `json:"name"`
		Author string `json:"author"`
	}

	test := Test {
		"This is a testa JSON",
		"Sondre Benjamin Aasen",
	}

	http.HandleFunc("/get/fake", internal.FakeTeamMembers)
	http.HandleFunc("/get/member", internal.GetTeamMembers)
	http.HandleFunc("/get/game", internal.GetGamesOfMember)
	http.HandleFunc("/get/connection", internal.GetTeamConnection)
	CO1Cache.Initialize()
	CO1Cache.WriteJSON("test", test)

	fmt.Println("Hello World!")

	http.ListenAndServe(":8080", nil)

}



