package rating

import (
	"log"
	"strconv"
	"time"
)

// Loops through every team member vs team member permutation and calls `getGamesOfMemberVSMember`.
func getGamesInTeam(teamMembers [] TeamMember) []Game {
	var games [] Game
	var tmpGames [] Game
	count := 0 // TODO REMOVE ALL INSTANCES OF count

	countOfMemberVSMember := 0
	for i := 0; i < len(teamMembers); i++ {
		countOfMemberVSMember += i+1
	}
	log.Println(countOfMemberVSMember)

	for i := 0; i < len(teamMembers); i++ {
		for j := i + 1; j < len(teamMembers); j++ {
			count++
			print("Match nr." + strconv.Itoa(count) + " of " + strconv.Itoa(countOfMemberVSMember))
			tmpGames = getGamesOfMemberVSMember(teamMembers[i], teamMembers[j])
			time.Sleep(5 * time.Second)
			for k := 0; k < len(tmpGames); k++ {
				games = append(games, tmpGames[k])
			}
		}
	}
	return games
}