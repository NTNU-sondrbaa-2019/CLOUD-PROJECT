
//TeamMembers - Struct for getting essential teammembers information
type TeamMembers struct {
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

func GetTeamMembers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var Teamids TeamMembers
			parts := strings.Split(r.URL.Path, "/")
			teamid := len(parts-1)
			resp, err := http.Get("https://lichess.org/team/" + teamid "/users")
			if err != nil {
				log.Fatalln(err)
				return
			}
			err = json.NewDecoder(response.Body).Decode(&Teamids)

			http.Header.Add(w.Header(), "content-type", "application/json")
			err = json.NewEncoder(w).Encode(Teamids)
			if err != nil {
				http.Error(w, "Could not encode json", 400)
				return
			}
		
		}
	}