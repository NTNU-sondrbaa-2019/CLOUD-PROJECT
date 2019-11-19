package internal

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

			for i = 0; i < 12; i++{
				print(teamMembers[i].ID + " " + teamMembers[i].Username + " " + strconv.Itoa(teamMembers[i].InternalCreatedAt))
			}

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
		request := "https://lichess.org/api/games/user/" + member.ID + "?vs=" + vsMember.ID +  "&perftype=blitz,classical,rapid,correspondence&since=" + strconv.Itoa(member.InternalCreatedAt)
		client := http.DefaultClient
		response := GetRequest(client, request)

		reader := bufio.NewReader(response.Body)
		var i = 0

		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		var tmp Game
		err = json.Unmarshal(line, &tmp)
		if err != nil {
			log.Print("Unmarshall Error:")
			log.Print(err)
		}
		games = append(games, tmp)
		i++

		for {
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
		}

		for i = 0; i < 1; i++{
			print(games[i].ID + ":\n\tWhite: " + games[i].Players.White.User.Name + "\n\tBlack: " + games[i].Players.Black.User.Name + "\n\tWinner: " + games[i].Winner)
		}
		return games
}
