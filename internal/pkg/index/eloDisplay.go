package index

import (
	"fmt"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/HTTPErrors"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/view"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type forTmp struct {
	LeagueName  string
	SeasonName  string
	Groups      []Group
	CurrentYear int
}

func EloDisplay(w http.ResponseWriter, r *http.Request) HTTPErrors.Error {
	season := r.URL.Query().Get("season")
	league, _ := strconv.ParseInt(r.URL.Query().Get("league"), 10, 64)

	fmt.Println(season, league)

	groups, err := database.SelectGroupsByLeagueIDAndLeagueSeasonName(league, season)
	if err != nil {
		view.ErrorPage(w, "Parameter Error", http.StatusInternalServerError)
	}

	selectedLeague, err := database.SelectLeagueSeason(league, season)
	if err != nil {
		view.ErrorPage(w, "Parameter Error", http.StatusInternalServerError)
	}

	var bigEloD forTmp

	for i, group := range *groups {

		fmt.Println(i, group)
		fmt.Println("Group ID: ", group.ID)

		result, err := database.SelectResultLastByGroupId(group.ID)

		if err == nil {

			var eloD Group

			eloD.Name = group.Name
			eloD.ID = group.ID
			eloD.LastPlayed = result.Played
			eloD.CurrentELO = result.ELOAfter

			bigEloD.Groups = append(bigEloD.Groups, eloD)

			fmt.Println("EloDisplay: ", eloD)

		}

	}

	bigEloD.LeagueName = selectedLeague.Name
	bigEloD.SeasonName = season
	bigEloD.CurrentYear = time.Now().Year()

	sort.Slice(bigEloD.Groups, func(i, j int) bool {
		return bigEloD.Groups[i].CurrentELO > bigEloD.Groups[j].CurrentELO
	})

	view.Render(w, "season_table", bigEloD)

	return HTTPErrors.NewError("", 0)
}
