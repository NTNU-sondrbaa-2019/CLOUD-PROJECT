package database

import (
	"encoding/json"
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

var connection *sqlx.DB

func Connect() {

	// TODO Fix logic

	if !CO1Cache.Verify("db-config") {

		log.Println("Writing default database configuration...")
		CO1Cache.WriteJSON("db-config", DEFAULT_CONNECTION)
		log.Fatalln("To test the AWS database from localhost, please insert password into the generated ./cache/db-config.json file!")

	} else {

		var db Connection
		var temp Connection
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

		temp.Database = os.Getenv("DB-DATABASE")

		if temp.Database != "" {
			db.Database = temp.Database
		}

		// TODO Sondre: Error handling of invalid HOST, PORT, USERNAME, PASSWORD, DATABASE

		if err != nil {
			log.Fatalln("Couldn't read database connection data...")
		}

		connection, err = sqlx.Open("mysql", db.Username+":"+db.Password+"@tcp("+db.Host+":"+db.Port+")/"+db.Database+"?parseTime=true")

		if err != nil {
			log.Fatalln("Couldn't connect to database: ", err)
		}
	}
}

func GetConnection() *sqlx.DB {
	return connection
}
