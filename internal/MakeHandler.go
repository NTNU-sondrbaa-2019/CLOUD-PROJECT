package internal

import (
	"fmt"
	"net/http"
)

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
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
