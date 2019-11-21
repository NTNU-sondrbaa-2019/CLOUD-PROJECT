package rating

import (
	"strconv"
	"time"
)

func GetGamesInTeam(teamMembers [] TeamMember) []Game{
	var games [] Game
	var tmpGames [] Game
	print("Third Member:" + teamMembers[2].Username + "\n")
	count := 0 // TODO REMOVE ALL INSTANCES OF count
	for i := 0; i < len(teamMembers); i++ {
		for j := i + 1; j < len(teamMembers); j++ {
			count++
			timeleft := (77 - count) * 5
			print("Match nr." + strconv.Itoa(count) + "\tTime remaining: " + strconv.Itoa(timeleft) + " ")
			tmpGames = GetGamesOfMemberVSMember(teamMembers[i], teamMembers[j])
			//print(tmpGames[i].Winner)
			//print(strconv.Itoa(i) + "  " +  strconv.Itoa(j) + "\t")
			//print(teamMembers[i].Username + " vs " + teamMembers[j].Username + "\n")
			time.Sleep(5 * time.Second)
			for k := 0; k < len(tmpGames); k++ {
				games = append(games, tmpGames[k])
			}
		}
	}

	for i := 0; i < len(games); i++ {
		print(games[i].Winner)
	}

	return games
}