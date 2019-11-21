package rating

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
)

func GetTeamElo(teamIdKey string) []TeamMember{
	var teamMembers [] TeamMember
	//parts := strings.Split(r.URL.Path, "/")
	//teamid := len(parts)-1
	request := "https://lichess.org/team/" + teamIdKey +  "/users"
	client := http.DefaultClient
	response := GetRequest(client, request)

	reader := bufio.NewReader(response.Body)
	var i = 0


	line, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
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

	//TODO Remove
	return CalculateElo(SortGames(GetGamesInTeam(teamMembers)), teamMembers)
}