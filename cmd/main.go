package main

import (
	"fmt"
	"github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/rating"
)

import "github.com/robfig/cron/v3"

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

	c := cron.New()
	teamIdKey := "storbukk-sjakklubb"
	_, err := c.AddFunc("0 2 * * *", func() {
		rating.GetTeamElo(teamIdKey)
	})

	if err != nil {
		panic(err)
	}

	c.Start()
}



