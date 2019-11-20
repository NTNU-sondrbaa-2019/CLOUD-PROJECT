package main

import (
	"fmt"
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/rating"
	"github.com/robfig/cron/v3"
	"net/http"
)

func main() {

	type Test struct {
		Name string `json:"name"`
		Author string `json:"author"`
	}

	test := Test {
		"This is a test JSON",
		"Sondre Benjamin Aasen",
	}



	CO1Cache.Initialize()
	CO1Cache.WriteJSON("test", test)

	fmt.Println("Hello World!")

	// Uncomment to run the lichess stuff.

	// Go service must be running for the cron job to take place
	c := cron.New()	
	teamIdKey := "storbukk-sjakklubb"
	//_, err := c.AddFunc("0 2 * * *", func() {
	// For testing purposes run every 10 minutes
	_, err := c.AddFunc("*/10 * * * *", func() {
		rating.GetTeamElo(teamIdKey)
	})

	if err != nil {
		panic(err)
	}

	c.Start()
	_ = http.ListenAndServe(":8080", nil)
}



