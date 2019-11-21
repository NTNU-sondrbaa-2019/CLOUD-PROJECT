package main

import (
	"encoding/json"
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/database"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/handler"

	"log"

	"os"

	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/rating"
	"github.com/robfig/cron/v3"
	"net/http"
	"strings"
)

func main() {

	// initialize cache
	log.Println("Initializing cache...")
	CO1Cache.Initialize()

	// TODO Sondre: Move to database package

	// write database config to cache
	log.Println("Initializing database...")
	if !CO1Cache.Verify("db-config") {
		log.Println("Writing default database configuration...")
		CO1Cache.WriteJSON("db-config", database.DEFAULT_CONNECTION)
		log.Fatalln("To test the AWS database from localhost, please insert password into the generated ./cache/db-config.json file!")
	} else {

		var db database.DatabaseConnection
		err := json.Unmarshal(CO1Cache.Read("db-config"), &db)

		db_host := os.Getenv("DB_HOST")

		if db_host != "" {
			db.Host = db_host
		}

		db_port := os.Getenv("DB_PORT")

		if db_port != "" {
			db.Port = db_port
		}

		db_username := os.Getenv("DB_USERNAME")

		if db_username != "" {
			db.Username = db_username
		}

		db_password := os.Getenv("DB_PASSWORD")

		if db_password != "" {
			db.Password = db_password
		}

		// TODO Sondre: Error handling of invalid HOST, PORT, USERNAME, PASSWORD

		if err != nil {
			log.Fatalln("Couldn't read database connection data...")
		}

		// TODO Sondre: connect to database
	}

	// Uncomment to run the lichess stuff.
	// Go service must be running for the cron job to take place
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
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
