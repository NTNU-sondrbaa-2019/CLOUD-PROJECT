package main

import (
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/rating"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/server"
	"log"
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
	rating.Initialize()

	// Starts the webserver
	server.Start()
}
