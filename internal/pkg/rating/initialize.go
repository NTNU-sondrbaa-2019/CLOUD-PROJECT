package rating

import (
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"strings"
)

func Initialize(){
	c := cron.New()

	// Get lichess.org teams to automatically fetch data about from environment
	_ = os.Setenv("LICHESS_TEAMS", "storbukk-sjakklubb,testclub")
	tmp := os.Getenv("LICHESS_TEAMS")
	teams := strings.Split(tmp, ",")

	if tmp != "" {
		log.Println("Checking elo for teams " + tmp)
		for i := 0; i < len(teams); i++ { // "0 2 * * *" every night 2am
			var team string
			team = teams[i]
			_, err := c.AddFunc("0 22 * * *", func() { // For testing use "min hour * * *" to set a time for the cronjob
				getTeamElo(team)
			})
			if err != nil {
				print(err)
			}
		}
	}else {		// If no teams, use storbukk-sjakklub
		// For testing purposes run every 10 minutes
		log.Println("No enviroment variable for teams, using default value '" + LICHESS_DEFAULT_TEAM + "'")
		_, err := c.AddFunc("0 22 * * *", func() {
			getTeamElo(LICHESS_DEFAULT_TEAM)
		})
		if err != nil {
			print(err)
		}
	}

	// Start cronjobs
	c.Start()

	log.Println("Initialized elo fetching!")
}
