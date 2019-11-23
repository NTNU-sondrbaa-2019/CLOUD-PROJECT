package api

import "github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"

// Struct to store data gathered
type returnType struct {
	TeamName string
	Users []database.USER
}

var someValues []returnType

var user *database.USER
var group *database.GROUP
var groups *[]database.GROUP
var group_user *[]database.GROUP_USER

