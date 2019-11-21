package rating

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
)

/*	GetTeamElo

	Parameter `teamIdKey` is the id of a team on lichess.org

	Does a GET request to lichess and gets all the members in a team. Calls `GetGamesInTeam` to get all games in a team.
	Calls `sortGames` to sort these games. Calls `calculateElo` to calculate Elo.

	returns `[]TeamMember` with all team members in a team with an internal ELO in team.
*/
func GetTeamElo(teamIDKey string) []TeamMember {
	var teamMembers [] TeamMember
	//parts := strings.Split(r.URL.Path, "/")
	//teamid := len(parts)-1
	request := "https://lichess.org/team/" + teamIDKey + "/users"
	client := http.DefaultClient
	response := getRequest(client, request)

	reader := bufio.NewReader(response.Body)
	var i = 0

	line, err := reader.ReadBytes('\n')
	if err != nil {
		log.Print(err)
	}
	var tmp TeamMember
	err = json.Unmarshal(line, &tmp)
	if err != nil {
		log.Print(err)
	}
	teamMembers = append(teamMembers, tmp)
	i++

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		var tmp TeamMember
		err = json.Unmarshal(line, &tmp)
		if err != nil {
			log.Print(err)
		}
		teamMembers = append(teamMembers, tmp)
		i++
	}

	return calculateElo(sortGames(getGamesInTeam(teamMembers)), teamMembers)
}
