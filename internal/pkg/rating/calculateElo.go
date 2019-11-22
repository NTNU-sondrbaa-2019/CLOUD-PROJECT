package rating

import (
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
	"math"
	"strconv"
)

func calculateElo(games []Game, teamMembers []TeamMember) []TeamMember {

	//TODO get elo from own database
	for i := 0; i < len(teamMembers); i++ {
		teamMembers[i].Elo = 1500.0
	}

	newTeamMembers := teamMembers

	var results [] database.RESULT
	// Loops through all games and calculates elo
	for i := 0; i < len(games); i++ {

		tmp := getWhiteAndBlackFromGame(games[i], newTeamMembers)
		white := tmp[0]
		black := tmp[1]

		var w float64
		var b float64

		// TODO START make this block a single function to calculate elo
		K := 64.0

		R1 := math.Pow(10, white.Elo/400)
		R2 := math.Pow(10, black.Elo/400)

		WhiteChance := R1 / (R1 + R2)
		BlackChance := R2 / (R1 + R2)

		if games[i].Winner == "white" {
			w = white.Elo + (K * (1 - WhiteChance))
			b = black.Elo + (K * (0 - BlackChance))
		} else if games[i].Winner == "black" {
			w = white.Elo + (K * (0 - WhiteChance))
			b = black.Elo + (K * (1 - BlackChance))
		} else {
			w = white.Elo + (K * (0.5 - WhiteChance))
			b = black.Elo + (K * (0.5 - BlackChance))
		}

		// TODO END

		newTeamMembers = insertElo(w, b, white, black, newTeamMembers)

		// TODO figure out how to create a database.RESULT object
		var result1 database.RESULT
		var result2 database.RESULT
		result1.GroupID = // DATABASE GROUP ID WTF DUDE
		result1.ELOBefore = white.Elo
		result2.ELOBefore = black.Elo


		result1.ELOAfter = w
		result1.Outcome = "" // win lose draw
		result1.Played = // time, unix convert
		result2.GroupID = // DATABASE GROUP ID WTF DUDE
		result1.ELOAfter = b
		result1.Outcome = ""
		result1.Played = // time, unix convert
		// TODO append to slice, for loop slice and insert into db

	}

	for i := 0; i < len(newTeamMembers); i++ {
		print("Elo for member " + newTeamMembers[i].Username + ":\t" + strconv.FormatFloat(newTeamMembers[i].Elo, 'f', -1, 64) + "\n")
	}

	return newTeamMembers
}
