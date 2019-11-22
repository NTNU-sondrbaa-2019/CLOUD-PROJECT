package main

import (
	"fmt"
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/gauth"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/handler"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/rating"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/team"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/user"
	"log"
	"net/http"
	"os"
)

import "github.com/robfig/cron/v3"

func main() {
	type Test struct {
		Name   string `json:"name"`
		Author string `json:"author"`
	}

	test := Test{
		"This is a test JSON",
		"Sondre Benjamin Aasen",
	}

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

	// 2 Options, either all handle requests are made here with new handlers
	// Or each endpoint can be handled via switch in a "mainHandler"
	//http.HandleFunc("/api/v1/", internal.MakeHandler(someHandler))

	http.HandleFunc("/", handler.MakeHandler(handler.HandleIndex))
	http.HandleFunc("/api/v1/", handler.MakeHandler(handler.HandleAPI))
	http.HandleFunc("/api/v1/team/", handler.MakeHandler(team.TeamHandler))
	http.HandleFunc("/api/v1/user/", handler.MakeHandler(user.UserHandler))
	http.HandleFunc("/api/v1/gauth/login/", handler.MakeHandler(gauth.LoginHandler))
	http.HandleFunc("/api/v1/gauth/loggedin/", handler.MakeHandler(gauth.LoggedInHandler))

	port := os.Getenv("PORT")

	if port == "" {
		port = handler.DEFAULT_PORT
	}

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
