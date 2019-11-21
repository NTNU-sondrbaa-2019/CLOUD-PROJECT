package rating

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetGamesOfMemberVSMember(member TeamMember, vsMember TeamMember) []Game{
	var games [] Game
	lastCreatedAt := 0
	if member.InternalCreatedAt < vsMember.InternalCreatedAt {
		lastCreatedAt = vsMember.InternalCreatedAt
	} else {
		lastCreatedAt = member.InternalCreatedAt
	}
	//TODO set lastCreatedAt to zero after we have made lastCreatedAt on a member
	//lastCreatedAt = 1572607209000 //11.01.2019
	lastCreatedAt = 1546350046000 //01.01.2019
	print( member.Username + "\t vs \t" + vsMember.Username + "\n")
	request := "https://lichess.org/api/games/user/" + member.Username + "?vs=" + vsMember.Username +  "&perftype=blitz,classical,rapid,correspondence&since=" + strconv.Itoa(lastCreatedAt)
	client := http.DefaultClient
	response := GetRequest(client, request)
	if response.StatusCode == 429 {
		log.Print("Rate limit on lichess.org reached. sleeping for 70 seconds...")
		time.Sleep(70 * time.Second)
		// TODO cronjob
		response = GetRequest(client, request)
	}
	reader := bufio.NewReader(response.Body)
	var i= 0

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
		i++
		ifPrint := false
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
			games = append(games, tmp)
			i++
			ifPrint = true
		}

		if ifPrint {
			for i = 0; i < 1; i++ {
				print("Game :\n\tWhite: " + games[i].Players.White.User.Name + "\n\tBlack: " + games[i].Players.Black.User.Name + "\n\tWinner: " + games[i].Winner + "\n")
			}
		}
	}
	return games
}