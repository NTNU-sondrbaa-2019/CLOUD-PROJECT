package database_test

import (
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
	"testing"
	"time"
)

func TestUser(t *testing.T) {

	_before(t)

	original := database.USER{}
	original.Name = "TestUser"
	original.Email = "test@localhost"
	original.Registered = time.Now()
	original.LastOnline = time.Now()

	id, err := database.InsertUser(original)

	if err != nil {
		_rollback()
		t.Fatal(err)
	}

	selected, err := database.SelectUserByID(*id)

	if err != nil {
		_rollback()
		t.Fatal(err)
	}

	if selected.Name != original.Name {
		t.Error("[INSERT] Names doesn't match.")
	}

	if selected.Email != original.Email {
		t.Error("[INSERT] Emails doesn't match.")
	}

	if selected.Registered.Unix() != original.Registered.Unix() {
		t.Errorf("[INSERT] Registered TIMEDATEs (%d and %d) doesn't match.", selected.Registered.Unix(), original.Registered.Unix())
	}

	if selected.LastOnline.Unix() != original.LastOnline.Unix() {
		t.Errorf("[INSERT] Last Online TIMEDATEs (%d and %d) doesn't match.", selected.LastOnline.Unix(), original.LastOnline.Unix())
	}

	modified := database.USER{}
	modified.Name = "NewUsername"

	err = database.ModifyUser(*id, modified)

	if err != nil {
		_rollback()
		t.Fatal(err)
	}

	selected, err = database.SelectUserByEmail(original.Email)

	if err != nil {
		_rollback()
		t.Fatal(err)
	}

	if selected.Name != modified.Name {
		t.Error("[MODIFY] Names doesn't match.")
	}

	if selected.Email != original.Email {
		t.Error("[MODIFY] Email doesn't match.")
	}

	if selected.Registered.Unix() != original.Registered.Unix() {
		t.Errorf("[MODIFY] Registered TIMEDATEs (%d and %d) doesn't match.", selected.Registered.Unix(), original.Registered.Unix())
	}

	if selected.LastOnline.Unix() != original.LastOnline.Unix() {
		t.Errorf("[MODIFY] Last Online TIMEDATEs (%d and %d) doesn't match.", selected.LastOnline.Unix(), original.LastOnline.Unix())
	}

	err = database.DeleteUser(*id)

	if err != nil {
		_rollback()
		t.Fatal(err)
	}

	users, err := database.SelectUsers("")

	if err != nil {
		_rollback()
		t.Fatal(err)
	}

	for _, user := range *users {
		if user.ID == *id {
			t.Error("[DELETE] User was not deleted.")
		}
	}

	_after(t)

}
