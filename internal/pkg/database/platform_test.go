package database_test

import (
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
	"testing"
)

func TestPlatform(t *testing.T) {

	CO1Cache.Initialize()
	database.Connect()

	_, err := database.GetConnection().Exec("START TRANSACTION")

	if err != nil {
		_, _ = database.GetConnection().Exec("ROLLBACK")
		t.Fatal(err)
	}

	// TODO tests

	_, _ = database.GetConnection().Exec("ROLLBACK")

}
