package main

import (
	"fmt"
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/handler"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/rating"
	"log"
	"net/http"
	"strings"
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
	// TODO accept input somewhere for the teamIdKey

	tmp := os.Getenv("LICHESS_TEAMS")
	teams := strings.Split(tmp, ",")


	if tmp != "" {
		for i := 0; i < len(teams); i++ {
			_, err := c.AddFunc("0 2 * * *", func() {
				rating.GetTeamElo(teams[i])
			})
			if err != nil {
				panic(err)
			}
		}
	}else {
		teamIdKey := "storbukk-sjakklubb"
		// For testing purposes run every 10 minutes
		_, err := c.AddFunc("*/10 * * * *", func() {
			rating.GetTeamElo(teamIdKey)
		})
		if err != nil {
			panic(err)
		}
	}

	c.Start()

	// 2 Options, either all handle requests are made here with new handlers
	// Or each endpoint can be handled via switch in a "mainHandler"
	//http.HandleFunc("/api/v1/", internal.MakeHandler(someHandler))

	http.HandleFunc("/", handler.MakeHandler(handler.HandleIndex))
	http.HandleFunc("/api/v1/", handler.MakeHandler(handler.HandleAPI))

	port := os.Getenv("PORT")

	if port == "" {
		port = handler.DEFAULT_PORT
	}

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
