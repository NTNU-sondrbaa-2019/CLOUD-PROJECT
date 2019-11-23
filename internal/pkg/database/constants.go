package database

type Connection struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

var DEFAULT_CONNECTION = Connection{
	"gr8elo.cekete5hvzfh.us-east-1.rds.amazonaws.com",
	"3306",
	"admin",
	"",
	"gr8elo",
}
