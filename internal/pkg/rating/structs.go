package rating

import "github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"

// TeamMember - Struct for getting essential teammembers information
type TeamMember struct {
	ID                string  `json:"id"`
	Username          string  `json:"username"`
	InternalCreatedAt int     `json:"internalCreatedAt"`
	Elo               float64 `json:"elo"`
	LeagueID          int64   `json:"league_id"`
}

// TODO check and fix struct for remi/draw
// Game - Struct for getting game/match information
type Game struct {
	ID        string          `json:"id"`
	CreatedAt int             `json:"createdAt"`
	Status    string          `json:"status"`
	Players   Players         `json:"players"`
	Winner    string          `json:"winner"`
	Result    database.RESULT `json:"result"`
}

type Players struct {
	White Player `json:"white"`
	Black Player `json:"black"`
}

type Player struct {
	User User `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
