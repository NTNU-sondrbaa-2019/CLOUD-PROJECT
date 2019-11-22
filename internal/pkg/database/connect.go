package database

import (
	"encoding/json"
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"
	"log"
	"os"
)

type DatabaseConnection struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var DEFAULT_CONNECTION = DatabaseConnection{
	"gr8elo.cekete5hvzfh.us-east-1.rds.amazonaws.com",
	"3306",
	"admin",
	"",
}

func Connect() {

	if !CO1Cache.Verify("db-config") {

		log.Println("Writing default database configuration...")
		CO1Cache.WriteJSON("db-config", DEFAULT_CONNECTION)
		log.Fatalln("To test the AWS database from localhost, please insert password into the generated ./cache/db-config.json file!")

	} else {

		var db DatabaseConnection
		var temp DatabaseConnection
		err := json.Unmarshal(CO1Cache.Read("db-config"), &db)

		temp.Host = os.Getenv("DB-HOST")

		if temp.Host != "" {
			db.Host = temp.Host
		}

		temp.Port = os.Getenv("DB-PORT")

		if temp.Port != "" {
			db.Port = temp.Port
		}

		temp.Username = os.Getenv("DB-USERNAME")

		if temp.Username != "" {
			db.Username = temp.Username
		}

		temp.Password = os.Getenv("DB-PASSWORD")

		if temp.Password != "" {
			db.Password = temp.Password
		}

		// TODO Sondre: Error handling of invalid HOST, PORT, USERNAME, PASSWORD

		if err != nil {

			log.Fatalln("Couldn't read database connection data...")

		}

		// TODO Sondre: connect to database
	}
}
