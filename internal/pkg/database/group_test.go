package database_test

import (
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/rating"
	"testing"
	"time"
)

func TestGroup(t *testing.T) {

	_before(t)

	inserted := database.GROUP{
		LeagueID:         rating.LICHESS_LEAGUE_ID,
		LeagueSeasonName: rating.LICHESS_LEAGUE_SEASON_NAME,
		Name:             "A Temporary Group Name",
		Created:          time.Now(),
	}

	t.Log("Inserting group...")
	_, err := database.InsertGroup(inserted)
	_error_fatal(t, err)
	t.Log("Group inserted!")

	_after(t)

}
