package rating

import (
	"testing"
)

func TestInsertElo(t *testing.T) {

	var PlayerAndElo []TeamMember
	PlayerAndElo = append(PlayerAndElo, TeamMember{Username: "Hyge"})
	PlayerAndElo = append(PlayerAndElo, TeamMember{Username: "andreas_nl"})
	PlayerAndElo = append(PlayerAndElo, TeamMember{Username: "Sondrir"})
	PlayerAndElo = append(PlayerAndElo, TeamMember{Username: "Sindrir"})
	PlayerAndElo = append(PlayerAndElo, TeamMember{Username: "Vebjornrir"})
	PlayerAndElo = append(PlayerAndElo, TeamMember{Username: "Klementin"})
	PlayerAndElo = append(PlayerAndElo, TeamMember{Username: "Carlsen"})

	WhiteElo := 2000.0
	BlackElo := 1600.0

	InsertedELO := insertElo(WhiteElo, BlackElo, PlayerAndElo[0], PlayerAndElo[1], PlayerAndElo)
	if InsertedELO[0].Elo != 2000.0 && InsertedELO[1].Elo != 1500.0 {
		t.Errorf("Unexpected value")
	}
}
