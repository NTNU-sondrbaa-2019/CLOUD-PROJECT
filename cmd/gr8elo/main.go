package main

import (
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/database"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/handler"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/rating"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	// initialize cache
	log.Println("Initializing cache...")
	CO1Cache.Initialize()

	// write database config to cache
	log.Println("Initializing database...")
	database.Connect()

	// Setting up cronjobs
	log.Println("Initializing automatic ELO fetching...")
	c := cron.New()

	// Get lichess.org teams to automatically fetch data about from environment
	lichess_teamids := os.Getenv("LICHESS-TEAMS")

	// Get default if nothing in environment
	if lichess_teamids == "" {
		lichess_teamids = rating.LICHESS_DEFAULT_TEAMS
	}

	// Add cronjob
	_, err := c.AddFunc("0 2 * * *", func() {
		for _, lichess_teamid := range strings.Split(lichess_teamids, ",") {
			rating.GetTeamElo(lichess_teamid)
		}
	})

	if err != nil {

		log.Println(err)

	}

	// Start cronjobs
	c.Start()

	// 2 Options, either all handle requests are made here with new handlers
	// Or each endpoint can be handled via switch in a "mainHandler"
	//http.HandleFunc("/api/v1/", internal.MakeHandler(someHandler))

	http.HandleFunc("/", handler.MakeHandler(handler.HandleIndex))
	http.HandleFunc("/api/v1/", handler.MakeHandler(handler.HandleAPI))
	http.HandleFunc("/api/v1/team/", handler.MakeHandler(handler.TeamHandler))

	port := os.Getenv("PORT")

	if port == "" {
		port = handler.DEFAULT_PORT
	}

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
