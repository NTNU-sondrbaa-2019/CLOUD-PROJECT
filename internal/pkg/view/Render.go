package view

import "net/http"

func Render(writer http.ResponseWriter, s string, page interface{}) {
	// Builds webpage and assigns data from page interface
	err := Templates.ExecuteTemplate(writer, s, page)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}