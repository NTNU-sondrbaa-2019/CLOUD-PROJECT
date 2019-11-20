package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^\\/(api\\/v1)\\/(\\S*)")

func main() {
	http.HandleFunc("/api/v1/", makeHandler(handlerTest))

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerTest(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Println("Shit still works")
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// m is the string after api/v1/ example "user/results/"
		m := validPath.FindStringSubmatch(r.URL.Path) // Checks if it matches criteria
		if m == nil {
			fmt.Println(m[3])
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2]) // Returns a "handler" with w(res) r(req) m(string)
	}
}
