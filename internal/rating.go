package internal

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type TeamMember struct {
	ID string `json:"id"`
	Username string `json:"username"`
	InternalCreatedAt int  `json:"internalCreatedAt"`
}

//TeamMember - Struct for getting essential teammembers information
type Game struct {
    ID        string `json:"id"`
    CreatedAt int  `json:"createdAt"`
    Players   struct {
        White struct {
            User struct {
                ID   string `json:"id"`
                Name string `json:"name"`
            } `json:"user"`
            Rating int `json:"rating"`
        } `json:"white"`
        Black struct {
            User struct {
                ID   string `json:"id"`
                Name string `json:"name"`
            } `json:"user"`
            Rating int `json:"rating"`
        } `json:"black"`
    } `json:"players"`
    Winner string `json:"winner"`
}

func FakeTeamMembers(w http.ResponseWriter, r *http.Request){
	var Teamids[12] TeamMember
	Teamids[0].ID = "iwindj"
	Teamids[1].ID = "JoakimPB"
	Teamids[2].ID = "BeagluZ"
	Teamids[3].ID = "NorTroll"
	Teamids[4].ID = "Atman96"
	Teamids[5].ID = "HermanDyrkorn"
	Teamids[6].ID = "Y3SH1"
	Teamids[7].ID = "TyrotoxismB"
	Teamids[8].ID = "noDiva"
	Teamids[9].ID = "StorbukkSK"
	Teamids[10].ID = "Hyge"
	Teamids[11].ID = "andreas_nl"

	Teamids[0].Username = "iwindj"
	Teamids[1].Username = "JoakimPB"
	Teamids[2].Username = "BeagluZ"
	Teamids[3].Username = "NorTroll"
	Teamids[4].Username = "Atman96"
	Teamids[5].Username = "HermanDyrkorn"
	Teamids[6].Username = "Y3SH1"
	Teamids[7].Username = "TyrotoxismB"
	Teamids[8].Username = "noDiva"
	Teamids[9].Username = "StorbukkSK"
	Teamids[10].Username = "Hyge"
	Teamids[11].Username = "andreas_nl"

	Teamids[0].InternalCreatedAt = 1572566400
	Teamids[1].InternalCreatedAt = 1572566400
	Teamids[2].InternalCreatedAt = 1572566400
	Teamids[3].InternalCreatedAt = 1572566400
	Teamids[4].InternalCreatedAt = 1572566400
	Teamids[5].InternalCreatedAt = 1572566400
	Teamids[6].InternalCreatedAt = 1572566400
	Teamids[7].InternalCreatedAt = 1572566400
	Teamids[8].InternalCreatedAt = 1572566400
	Teamids[9].InternalCreatedAt = 1572566400
	Teamids[10].InternalCreatedAt = 1572566400
	Teamids[11].InternalCreatedAt = 1572566400



	http.Header.Add(w.Header(), "content-type", "application/json")
	err := json.NewEncoder(w).Encode(Teamids)
	if err != nil {
		http.Error(w, "Could not encode json", 400)
		return
	}
}

func GetRequest(c *http.Client, s string) *http.Response {
	req, err := http.NewRequest("GET", s, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Accept", "application/x-ndjson")
	resp, err := c.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}

func GetTeamMembers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			var teamMembers [] TeamMember
			//parts := strings.Split(r.URL.Path, "/")
			//teamid := len(parts)-1
			teamidkey := r.URL.Query().Get("teamid")
			request := "https://lichess.org/team/" + teamidkey +  "/users"
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
			GetGamesInTeam(teamMembers)

			http.Header.Add(w.Header(), "content-type", "application/json")
			err = json.NewEncoder(w).Encode(teamMembers)
			if err != nil {
				http.Error(w, "Could not encode json", 400)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
}

func GetGamesOfMember(member TeamMember, vsMember TeamMember) []Game{
		var games [] Game
		lastCreatedAt := 0
		if member.InternalCreatedAt < vsMember.InternalCreatedAt {
			lastCreatedAt = vsMember.InternalCreatedAt
		} else {
			lastCreatedAt = member.InternalCreatedAt
		}
		//TODO set lastCreatedAt to zero after we have made lastCreatedAt on a member
		lastCreatedAt = 1572607209000
		print("Match: " + member.Username + "\t vs \t" + vsMember.Username + "\n")
		request := "https://lichess.org/api/games/user/" + member.Username + "?vs=" + vsMember.Username +  "&perftype=blitz,classical,rapid,correspondence&since=" + strconv.Itoa(lastCreatedAt)
		client := http.DefaultClient
		response := GetRequest(client, request)
		if response.StatusCode == 429 {
			log.Print("Rate limit on lichess.org reached. sleeping for 70 seconds...")
			time.Sleep(70 * time.Second)
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

func GetGamesInTeam(teamMembers [] TeamMember) []Game{
	var games [] Game
	var tmpGames [] Game
	print("Third Member:" + teamMembers[2].Username + "\n")

	for i := 0; i < len(teamMembers); i++ {
		for j := i + 1; j < len(teamMembers); j++ {
			tmpGames = GetGamesOfMember(teamMembers[i], teamMembers[j])
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