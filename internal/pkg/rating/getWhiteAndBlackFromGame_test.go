package rating

import (
	"testing"
)

func TestGetWhiteAndBlackFromGame(t *testing.T) {

	var games []Game
	games = append(games, Game{Players: Players{White: Player{User: User{Name: "Hyge"}}, Black: Player{User{Name: "andreas_nl"}}}})
	games = append(games, Game{Players: Players{White: Player{User: User{Name: "Sindrir"}}, Black: Player{User{Name: "Carlsen"}}}})

	var teamMember []TeamMember
	teamMember = append(teamMember, TeamMember{Username: "Hyge"})
	teamMember = append(teamMember, TeamMember{Username: "andreas_nl"})
	teamMember = append(teamMember, TeamMember{Username: "Sindrir"})
	teamMember = append(teamMember, TeamMember{Username: "Carlsen"})

	var TeammemberExpected []TeamMember
	TeammemberExpected = append(TeammemberExpected, TeamMember{Username: "Hyge"})
	TeammemberExpected = append(TeammemberExpected, TeamMember{Username: "andreas_nl"})
	TeammemberExpected = append(TeammemberExpected, TeamMember{Username: "Sindrir"})
	TeammemberExpected = append(TeammemberExpected, TeamMember{Username: "Carlsen"})

	WhiteAndBlackFromGame := getWhiteAndBlackFromGame(games[0], teamMember)

	for i := 0; i < len(WhiteAndBlackFromGame); i++ {
		if WhiteAndBlackFromGame[0].Username != TeammemberExpected[0].Username &&
			WhiteAndBlackFromGame[1].Username != TeammemberExpected[1].Username {
			t.Errorf("Unexpected value")
		}
	}
}
