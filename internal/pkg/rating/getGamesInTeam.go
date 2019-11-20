package rating

import (
	"strconv"
	"time"
)

func getGamesInTeam(teamMembers [] TeamMember) []Game {
	var games [] Game
	var tmpGames [] Game
	count := 0 // TODO REMOVE ALL INSTANCES OF count
	for i := 0; i < len(teamMembers); i++ {
		for j := i + 1; j < len(teamMembers); j++ {
			count++
			timeleft := (77 - count) * 5
			print("Match nr." + strconv.Itoa(count) + "\tTime remaining: " + strconv.Itoa(timeleft) + " ")
			tmpGames = getGamesOfMemberVSMember(teamMembers[i], teamMembers[j])
			//print(tmpGames[i].Winner)
			//print(strconv.Itoa(i) + "  " +  strconv.Itoa(j) + "\t")
			//print(teamMembers[i].Username + " vs " + teamMembers[j].Username + "\n")
			time.Sleep(5 * time.Second)
			for k := 0; k < len(tmpGames); k++ {
				games = append(games, tmpGames[k])
			}
		}
	}
	return games
}