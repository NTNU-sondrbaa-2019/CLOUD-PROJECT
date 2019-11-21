package database

type DatabaseConnection struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var DEFAULT_CONNECTION = DatabaseConnection{
	"gr8elo.cekete5hvzfh.us-east-1.rds.amazonaws.com",
	"3306",
	"admin",
	"",
}

func Connect() {

}
