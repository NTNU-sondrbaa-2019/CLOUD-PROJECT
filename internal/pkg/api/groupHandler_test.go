package api

import (
	"fmt"
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/HTTPErrors"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestGroupHandler(t *testing.T) {
	var groups *[]database.GROUP


	CO1Cache.Initialize()
	database.Connect()

	_, err := database.GetConnection().Exec("START TRANSACTION")

	if err != nil {
		_, _ = database.GetConnection().Exec("ROLLBACK")
		t.Fatal(err)
	}

	// Defining Test Data
	var testGroup database.GROUP

	testGroup.Name = "TestName"
	layout := "2006-01-02T15:04:05.000Z"
	testGroup.Created, err = time.Parse(layout, "2014-11-12T15:00:00.371Z")
	if err != nil {
		t.Fatal(err)
	}
	testGroup.LeagueID = 3
	testGroup.LeagueSeasonName = "Sjakksesong"

	var testUser database.USER

	testUser.Name = "TestUser"
	testUser.LastOnline, err = time.Parse(layout, "2014-11-12T15:00:00.371Z")
	if err != nil {
		t.Fatal(err)
	}
	testUser.Registered, err = time.Parse(layout, "2014-11-12T15:00:00.371Z")
	if err != nil {
		t.Fatal(err)
	}
	testUser.Email = "test@test.com"


	// Inserting Test Data to Database
	database.InsertGroup(testGroup)
	database.InsertUser(testUser)

	var returnedGroup database.GROUP
	// Need to get id's generated from database to delete
	groups, err = database.SelectGroups("WHERE name=\"" + testGroup.Name + "\"")
	if err != nil {
		t.Fatal(err)
	}
	for i, g := range *groups {
		fmt.Println(i,g) // For debugging, prints i and group
		 returnedGroup = g
	}
	testGroup.ID = returnedGroup.ID
	testGroup.Created = returnedGroup.Created

	returnedUser, _ := database.SelectUserByEmail(testUser.Email)
	testUser.ID = returnedUser.ID
	testUser.Registered = returnedUser.Registered

	// Setting test data for GROUP_USER
	var testGroupUser database.GROUP_USER
	testGroupUser.GroupID = testGroup.ID
	testGroupUser.UserID = testUser.ID

	database.InsertGroupUser(testGroupUser)
	//var returnedGroupUser, _ = database.SelectGroupUserByGroupID(testGroupUser.GroupID)

	// Returned database structs
	// returnedUser
	// returnedGroup
	// returnedGroupUser

	// Testing api/v1/group/testGroup
	req, err := http.NewRequest("GET", "/api/v1/team/" + testGroup.Name, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler:= makeHandler(GroupHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	fmt.Println(returnedGroup)
	fmt.Println(rr.Body)
	// Check the response body is what we expect.
	expected := `[{"id":"` + strconv.FormatInt(testGroup.ID,10) + `","league_id":"` + strconv.FormatInt(testGroup.LeagueID,10) + `","league_season_name":"` + testGroup.LeagueSeasonName + `","name":"`+testGroup.Name+`","created":"2019-11-22 14:16:06"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	// Deleting Test Data
	database.DeleteGroup(returnedGroup.ID)
	database.DeleteGroupUser(returnedGroup.ID, returnedUser.ID)
	database.DeleteUser(returnedUser.ID)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request) HTTPErrors.Error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
