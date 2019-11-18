package internal

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
	"net/http"
)

type TeamMembers struct {
	ID string `json:"id"`
	Username string `json:"username"`
	CreatedAt int  `json:"createdAt"`
}

//TeamMembers - Struct for getting essential teammembers information
type TeamGames struct {
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
	var Teamids[12] TeamMembers
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

	Teamids[0].CreatedAt = 1572566400
	Teamids[1].CreatedAt = 1572566400
	Teamids[2].CreatedAt = 1572566400
	Teamids[3].CreatedAt = 1572566400
	Teamids[4].CreatedAt = 1572566400
	Teamids[5].CreatedAt = 1572566400
	Teamids[6].CreatedAt = 1572566400
	Teamids[7].CreatedAt = 1572566400
	Teamids[8].CreatedAt = 1572566400
	Teamids[9].CreatedAt = 1572566400
	Teamids[10].CreatedAt = 1572566400
	Teamids[11].CreatedAt = 1572566400



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
	resp, err := c.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}

func GetTeamConnection(w http.ResponseWriter, r *http.Request) {
	teamidkey := r.URL.Query().Get("teamid")
	conn, err := net.Dial("tcp", "lichess.org:9000")
	request := "https://lichess.org/team/" + teamidkey +  "/users"
	client := http.DefaultClient
	response := GetRequest(client, request)
	log.Print(response.Header)

	if err != nil {
		log.Print(err)
	}

	status, err := bufio.NewReader(conn).ReadString('\n')
	log.Print(status)
	http.Header.Add(w.Header(), "content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func GetTeamMembers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var Teamids[] TeamMembers
		//parts := strings.Split(r.URL.Path, "/")
		//teamid := len(parts)-1
		teamidkey := r.URL.Query().Get("teamid")
		request := "https://lichess.org/team/" + teamidkey +  "/users"
		client := http.DefaultClient
		response := GetRequest(client, request)
		log.Print(response)


		reader := bufio.NewReader(response.Body)
		var i = 0
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(line))
		var tmp TeamMembers
		json.Unmarshal(line, &tmp)
		Teamids = append(Teamids, tmp)
		i++


		for string(line) != "EOF" {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				continue
			}


			var tmp TeamMembers
			json.Unmarshal(line, &tmp)
			log.Print(tmp.Username)
			Teamids = append(Teamids, tmp)
			i++

		}

		for i = 0; i < 12; i++{
			print(Teamids[i].Username)
		}




		//line, err := resp.Read()

		//log.Print(body)

		/*
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		err1 := json.Unmarshal(body, &Teamids)
		if err1 != nil {
			print("UnmarshallError: ")
			log.Print(err1)
		}

		http.Header.Add(w.Header(), "content-type", "application/json")
		err = json.NewEncoder(w).Encode(Teamids)
		if err != nil {
			http.Error(w, "Could not encode json", 400)
			return
		}

		 */

		}
	}


