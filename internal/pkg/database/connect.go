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

	var db Connection
	var err error

	db.Host = os.Getenv("DB-HOST")
	db.Port = os.Getenv("DB-PORT")
	db.Username = os.Getenv("DB-USERNAME")
	db.Password = os.Getenv("DB-PASSWORD")
	db.Database = os.Getenv("DB-DATABASE")

	dsn := db.Username + ":" + db.Password + "@tcp(" + db.Host + ":" + db.Port + ")/" + db.Database + "?parseTime=true"
	connection, err = sqlx.Open("mysql", dsn)

	if err != nil || connection.Ping() != nil {

		log.Println("Couldn't connect to database using environment variables: ", err)

		if !CO1Cache.Verify("db-config") {

			log.Println("Writing default database configuration...")
			CO1Cache.WriteJSON("db-config", DEFAULT_CONNECTION)
			log.Fatalln("To test the AWS database from localhost, please insert password into the generated ./cache/db-config.json file!")

		} else {

			err = json.Unmarshal(CO1Cache.Read("db-config"), &db)

			dsn := db.Username + ":" + db.Password + "@tcp(" + db.Host + ":" + db.Port + ")/" + db.Database + "?parseTime=true"
			connection, err = sqlx.Open("mysql", dsn)

			if err != nil {
				log.Fatalln("Couldn't connect to database using .cache/db-config.json configuration: ", err)
			}

		}

	}

	defer connection.Close()

	log.Println("Successfully connected to database!")
}

func GetConnection() *sqlx.DB {
	return connection
}
