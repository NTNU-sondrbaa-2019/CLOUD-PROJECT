package handler

import (
	"net/http"
	"path/filepath"
)

func HandleIndex(w http.ResponseWriter, r *http.Request, url string) {
	if url != "/" {
		http.NotFound(w, r)
	} else {
		lp := filepath.Join("templates", "web/view.html")
		fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

		tmpl, _ := template.ParseFiles(lp, fp)
		tmpl.ExecuteTemplate(w, "layout", nil)
	}
}

