package database_test

import (
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
	"testing"
)

func _before(t *testing.T) {

	CO1Cache.Initialize()
	database.Connect()

	_, err := database.GetConnection().Exec("START TRANSACTION")

	if err != nil {
		_rollback()
		t.Fatal(err)
	}

}

func _after(t *testing.T) {

	_rollback()

}

func _rollback() {

	_, _ = database.GetConnection().Exec("ROLLBACK")

}

func _error_fatal(t *testing.T, err error) {
	if err != nil {
		_rollback()
		t.Fatal(err)
	}
}

func _error_error(t *testing.T, err error) {
	if err != nil {
		_rollback()
		t.Error(err)
	}
}
