package rating

import (
	"log"
	"net/http"
)

func getRequest(c *http.Client, s string) *http.Response {
	req, err := http.NewRequest("GET", s, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Accept", "application/x-ndjson")
	resp, err := c.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}
