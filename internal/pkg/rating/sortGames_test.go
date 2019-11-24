package rating

import (
	"testing"
)

//1325376000 (Time since lichess was created)

func TestSortGames(t *testing.T) {

	var games []Game
	games = append(games, Game{CreatedAt: 1574431087810})
	games = append(games, Game{CreatedAt: 1475361788845})
	games = append(games, Game{CreatedAt: 1523244555843})
	games = append(games, Game{CreatedAt: 1444571985832})
	games = append(games, Game{CreatedAt: 1556471895843})
	games = append(games, Game{CreatedAt: 1545461298353})
	games = append(games, Game{CreatedAt: 1546679897863})

	var gamesExpected []Game
	gamesExpected = append(games, Game{CreatedAt: 1444571985832})
	gamesExpected = append(games, Game{CreatedAt: 1475361788845})
	gamesExpected = append(games, Game{CreatedAt: 1523244555843})
	gamesExpected = append(games, Game{CreatedAt: 1545461298353})
	gamesExpected = append(games, Game{CreatedAt: 1546679897863})
	gamesExpected = append(games, Game{CreatedAt: 1556471895843})
	gamesExpected = append(games, Game{CreatedAt: 1574431087810})

	gamesFromFunction := sortGames(games)

	for i := 0; i < len(gamesFromFunction); i++ {
		if gamesFromFunction[i].CreatedAt != gamesExpected[i].CreatedAt {
			t.Errorf("Unexpected value")
		}
	}

}
