package rating

import (
	"github.com/robfig/cron/v3"
	"os"
	"strings"
)

func Initialize(){
	c := cron.New()

	// Get lichess.org teams to automatically fetch data about from environment
	_ = os.Setenv("LICHESS_TEAMS", "storbukk-sjakklubb,testclub")
	tmp := os.Getenv("LICHESS-TEAMS")
	teams := strings.Split(tmp, ",")

	if tmp != "" {
		for i := 0; i < len(teams); i++ { // "0 2 * * *" every night 2am
			var team string
			team = teams[i]
			_, err := c.AddFunc("*/10 * * * *", func() {
				GetTeamElo(team)
			})
			if err != nil {
				print(err)
			}
		}
	}else {		// If no teams, use storbukk-sjakklub
		// For testing purposes run every 10 minutes
		_, err := c.AddFunc("*/10 * * * *", func() {
			GetTeamElo(LICHESS_DEFAULT_TEAMS)
		})
		if err != nil {
			panic(err)
		}
	}

	// Start cronjobs
	c.Start()
}
