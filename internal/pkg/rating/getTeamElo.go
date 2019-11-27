package rating

import (
	"bufio"
	"encoding/json"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
	"log"
	"net/http"
	"time"
)

/*	getTeamElo

	Parameter `teamIdKey` is the id of a team on lichess.org

	Does a GET request to lichess and gets all the members in a team. Calls `GetGamesInTeam` to get all games in a team.
	Calls `sortGames` to sort these games. Calls `calculateElo` to calculate Elo.

	returns `[]TeamMember` with all team members in a team with an internal ELO in team.
*/
func getTeamElo(teamIDKey string) []TeamMember {
	var teamMembers [] TeamMember
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

	// TODO CHECK IF LEAGUE/TEAM EXISTS IN DATABASE
	// IF NOT INSERT LEAGUE/TEAM AND MAKE A SEASON
	// INSERT ALL MEMBERS OF THE TEAM AS GROUPS

	// GROUP IS MEMBERS THAT PLAYS IN A SPECIFIC SEASON

	// TODO CHECK IF MEMBER EXISTS AS A GROUP IN THE LEAGUE/SEASON?
	// IF NOT INSERT MEMBER IN THE LEAGUE AS GROUP

	var league database.LEAGUE
	var leagueSeason database.LEAGUE_SEASON
	league.Name = teamIDKey
	league.PlatformEloID = 1
	exists, err := database.LeagueExists(league.Name)
	if err == nil {
		if exists {
			_, err = database.InsertLeague(league)
			if err == nil {
				leagueSeason.Name = "Sjakksesong"
				leagueSeason.LeagueID = 1

				_, err = database.InsertLeagueSeason(leagueSeason)
				if err != nil {
					log.Println(err)
				}
			} else {
				log.Println(err)
			}
		}
	}

	for i := 0; i < len(teamMembers); i++ {
		var tmpGroup database.GROUP

		leageIDFromDB, _ := database.SelectLeagueByPlatformEloIDAndName(1, teamIDKey)
		tmpGroup.LeagueID = leageIDFromDB.ID // TODO GET FROM DB

		tmpGroup.Name = teamMembers[i].Username
		tmpGroup.LeagueSeasonName = "Sjakksesong"
		tmpGroup.Created = time.Now()
		teamMembers[i].LeagueID = leageIDFromDB.ID
		_, err = database.InsertGroup(tmpGroup)
		if err != nil {
			log.Println(err)
		}
		log.Println("Inserting member " + teamMembers[i].Username + " into the db")
	}

	return calculateElo(sortGames(getGamesInTeam(teamMembers)), teamMembers)
}
