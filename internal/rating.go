package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type TeamMembers struct {
	ID string `json:"id"`
	Username string `json:"username"`
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

	http.Header.Add(w.Header(), "content-type", "application/json")
	err := json.NewEncoder(w).Encode(Teamids)
	if err != nil {
		http.Error(w, "Could not encode json", 400)
		return
	}
}

func GetTeamMembers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var Teamids TeamMembers
			//parts := strings.Split(r.URL.Path, "/")
			//teamid := len(parts)-1
			teamidkey := r.URL.Query().Get("teamid")
			resp, err := http.Get("https://lichess.org/team/" + teamidkey +  "/users")
			if err != nil {
				log.Fatalln(err)
				return
			}


			//err = json.NewDecoder(resp.Body).Decode(&Teamids.Members)

			/*if err != nil {
				log.Print(err)
			}
*/
			body, err := ioutil.ReadAll(resp.Body)

			//log.Print(body)
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

		}
	}


