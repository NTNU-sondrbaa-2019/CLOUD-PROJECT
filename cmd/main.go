package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"

	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/gauth"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/root"
)

func main() {
	// Initialize the local cache
	CO1Cache.Initialize()

	// Get the port, or set it to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", root.NilHandler)
	http.HandleFunc("/login", gauth.LoginHandler)
	http.HandleFunc("/logout", gauth.LogoutHandler)
	http.HandleFunc("/loggedin", gauth.LoggedInHandler)
	http.HandleFunc("/oauth2callback", gauth.OauthCallBackHandler)

	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
