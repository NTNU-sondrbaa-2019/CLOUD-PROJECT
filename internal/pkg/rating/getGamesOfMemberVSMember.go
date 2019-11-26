package rating

import (
	"bufio"
	"encoding/json"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Gets all matches between two members.
func getGamesOfMemberVSMember(member TeamMember, vsMember TeamMember) []Game {
	var games [] Game
	/*lastCreatedAt := 0
	if member.InternalCreatedAt < vsMember.InternalCreatedAt {
		lastCreatedAt = vsMember.InternalCreatedAt
	} else {
		lastCreatedAt = member.InternalCreatedAt
	}

	 */

	// todo this since variable does not make sense. Can make other teams games disapear
	since, err := database.SelectResultLastPlayedByLeagueID(member.LeagueID)
	var sinceTime int
	if err == nil && since != nil {
		sinceTime = (int(since.UnixNano())/1000000) + 1000
		log.Println("Unix epoch time in milliseconds:\t", sinceTime)
	}  else {
		log.Println(err)
		// TODO set to leagues start time
		sinceTime = 1572607209000 // 01.11.2019
	}
	print(member.Username + "\t vs \t" + vsMember.Username + "\n")
	request := "https://lichess.org/api/games/user/" + member.Username + "?vs=" + vsMember.Username + "&perftype=blitz,classical,rapid,correspondence&since=" + strconv.Itoa(sinceTime)
	client := http.DefaultClient
	response := getRequest(client, request)
	if response.StatusCode == 429 {
		log.Print("Rate limit on lichess.org reached. sleeping for " + strconv.Itoa(LICHESS_RATE_LIMIT) + " seconds...")
		time.Sleep(LICHESS_RATE_LIMIT * time.Second) // Waiting over 1 min for lichess' rate limit
		response = getRequest(client, request)
	}

	// START of code needed to parse the ndjson

	reader := bufio.NewReader(response.Body)

	line, err := reader.ReadBytes('\n')
	if string(line) != "" {
		if err != nil {
			log.Print(err)
		}
		var tmp Game
		err = json.Unmarshal(line, &tmp)
		if err != nil {
			log.Print("Unmarshall Error:")
			log.Print(err)
		}
		games = append(games, tmp)
		for {
			if err != nil {
				break
			}
			line, err := reader.ReadBytes('\n')
			if err != nil {
				break
			}
			print(line)
			var tmp Game
			err = json.Unmarshal(line, &tmp)
			if err != nil {
				log.Print("Unmarshall Error:")
				log.Print(err)
			}
			//games = append(games, Game{"",1,"", {{{"Hyge", "Hyge"}}, {{"Hyge", "Hyge"}}}, "black"})
			games = append(games, tmp)
		}

	}

		// END of code needed to parse the ndjson
	return games
}
