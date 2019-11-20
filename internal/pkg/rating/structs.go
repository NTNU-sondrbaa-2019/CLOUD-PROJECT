package rating

//TeamMember - Struct for getting essential teammembers information
type TeamMember struct {
	ID                string  `json:"id"`
	Username          string  `json:"username"`
	InternalCreatedAt int     `json:"internalCreatedAt"`
	Elo               float64 `json:"elo"`
}

type Game struct {
	ID        string `json:"id"`
	CreatedAt int    `json:"createdAt"`
	Players   struct {
		White struct {
			User struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"user"`
		} `json:"white"`
		Black struct {
			User struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"user"`
		} `json:"black"`
	} `json:"players"`
	Winner string `json:"winner"`
}
