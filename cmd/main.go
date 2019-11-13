package main

import (
	"fmt"
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"
	"log"
	"net/http"
	"os"

	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/root"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/gauth"
)

func main() {

	/*
	type Test struct {
		Name string `json:"name"`
		Author string `json:"author"`
	}

	test := Test {
		"This is a test JSON",
		"Sondre Benjamin Aasen",
	}

	CO1Cache.Initialize()
	CO1Cache.WriteJSON("test", test)

	fmt.Println("Hello World!")
	*/

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", root.NilHandler)
	http.HandleFunc("/login", gauth.LoginHandler)
	// http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/oauth2callback", gauth.OauthCallBackHandler)

	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
