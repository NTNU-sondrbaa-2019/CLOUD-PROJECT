package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"
<<<<<<< HEAD

	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/gauth"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/root"
=======
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/rating"
>>>>>>> 1429692dc26a9c8de6e4764cdb3d57f14a9b20ed
)

import "github.com/robfig/cron/v3"

func main() {
	// Initialize the local cache
	CO1Cache.Initialize()

	// Get the port, or set it to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

<<<<<<< HEAD
	http.HandleFunc("/", root.NilHandler)
	http.HandleFunc("/login", gauth.LoginHandler)
	http.HandleFunc("/logout", gauth.LogoutHandler)
	http.HandleFunc("/loggedin", gauth.LoggedInHandler)
	http.HandleFunc("/oauth2callback", gauth.OauthCallBackHandler)

	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
=======


	CO1Cache.Initialize()
	CO1Cache.WriteJSON("test", test)

	fmt.Println("Hello World!")

	// Uncomment to run the lichess stuff.

	c := cron.New()
	teamIdKey := "storbukk-sjakklubb"
	_, err := c.AddFunc("0 2 * * *", func() {
		rating.GetTeamElo(teamIdKey)
	})

	if err != nil {
		panic(err)
	}

	c.Start()
>>>>>>> 1429692dc26a9c8de6e4764cdb3d57f14a9b20ed
}



