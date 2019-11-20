package internal

import (
	"net/http"
	"regexp"
)

// Limits valid paths to host/api/v1/something
var validPath = regexp.MustCompile("^\\/(api\\/v1)\\/(\\S*)")

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// m is the string after api/v1/ example "user/results/"
		m := validPath.FindStringSubmatch(r.URL.Path) // Checks if it matches criteria
		if m == nil {
			http.NotFound(w, r)
			return
		}
		// m[2] is the second grouping from the regexp
		fn(w, r, m[2]) // Returns a "handler" with w(res) r(req) m(string)
	}
}
