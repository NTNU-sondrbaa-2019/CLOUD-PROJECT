package database_test

import (
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
	"testing"
	"time"
)

func TestUser(t *testing.T) {

	_before(t)

	inserted := database.USER{
		Name:       "TestUser",
		Email:      "test@localhost",
		Registered: time.Now(),
		LastOnline: time.Now(),
	}

	modified := database.USER{Name: "NewUsername"}

	t.Log("Inserting user...")
	user_id, err := database.InsertUser(inserted)
	_error_fatal(t, err)
	t.Log("User inserted!")

	t.Log("Selecting user by ID...")
	selected, err := database.SelectUserByID(*user_id)
	_error_fatal(t, err)
	t.Log("User selected!")
	t.Log("Comparing inserted user with selected user:")

	if selected.Name != inserted.Name {
		t.Error("\t- Names doesn't match.")
	} else {
		t.Log("\t- Names match.")
	}

	if selected.Email != inserted.Email {
		t.Error("\t- Emails doesn't match.")
	} else {
		t.Log("\t- Emails match.")
	}

	if selected.Registered.Unix() != inserted.Registered.Unix() {
		t.Errorf("\t- Registered timedates doesn't match.")
	} else {
		t.Log("\t- Registered timedates match.")
	}

	if selected.LastOnline.Unix() != inserted.LastOnline.Unix() {
		t.Errorf("\t- Last Online timedates doesn't match.")
	} else {
		t.Log("\t- Last Online timedates match.")
	}

	t.Log("Modifying user...")
	err = database.ModifyUser(*user_id, modified)
	_error_fatal(t, err)
	t.Log("User modified!")

	t.Log("Selecting user by email...")
	selected, err = database.SelectUserByEmail(inserted.Email)
	_error_fatal(t, err)
	t.Log("User selected!")
	t.Log("Comparing modified user with selected user:")

	if selected.Name != modified.Name {
		t.Error("\t- Names doesn't match.")
	} else {
		t.Log("\t- Names match.")
	}

	if selected.Email != inserted.Email {
		t.Error("\t- Emails doesn't match.")
	} else {
		t.Log("\t- Emails match.")
	}

	if selected.Registered.Unix() != inserted.Registered.Unix() {
		t.Errorf("\t- Registered timedates doesn't match.")
	} else {
		t.Log("\t- Registered timedates match.")
	}

	if selected.LastOnline.Unix() != inserted.LastOnline.Unix() {
		t.Errorf("\t- Last Online timedates doesn't match.")
	} else {
		t.Log("\t- Last Online timedates match.")
	}

	t.Log("Deleting user...")
	err = database.DeleteUser(*user_id)
	_error_fatal(t, err)
	t.Log("User deleted!")

	t.Log("Selecting all users...")
	users, err := database.SelectUsers("")
	_error_fatal(t, err)
	t.Log("Users selected!")

	user_deleted := true

	t.Log("Checking if user was deleted...")
	for _, user := range *users {
		if user.ID == *user_id {
			user_deleted = false
			t.Error("User was not deleted.")
		}
	}

	if user_deleted {
		t.Log("User was deleted correctly!")
	}

	_after(t)

}
