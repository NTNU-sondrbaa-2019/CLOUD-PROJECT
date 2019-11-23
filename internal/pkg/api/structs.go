package api

import "github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"

// Struct to store data gathered
type teamsUser struct {
	TeamName string `json:"groupname"`
	Users []database.USER `json:"members"`
}

type userTeams struct {
	Username string `json:"username"`
	UserID int64	`json:"user_id"`
	Groups []database.GROUP `json:"groups"`
}

type groupRes struct {
	GroupName string             `json:"groupname"`
	Results   *[]database.RESULT `json:"results"`
}

var someValues []teamsUser

var user *database.USER
var group *database.GROUP
var groups *[]database.GROUP
var group_user *[]database.GROUP_USER
var result *[]database.RESULT

