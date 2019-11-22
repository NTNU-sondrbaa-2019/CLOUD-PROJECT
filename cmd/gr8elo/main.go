package main

import (
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/rating"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/server"
	"github.com/robfig/cron/v3"
	"log"
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

	// Starts the webserver
	server.Start()
}
