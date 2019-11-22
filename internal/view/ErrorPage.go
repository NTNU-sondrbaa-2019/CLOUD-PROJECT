package view

import (
	"net/http"
	"strconv"
	"time"
)

func ErrorPage(w http.ResponseWriter,errorMsg string, code int) {
	type err struct {
		ErrorMsg string
		ErrorCode int
		CurrentYear string
	}
	error := &err{ErrorMsg: errorMsg, ErrorCode: code, CurrentYear: strconv.Itoa(time.Now().Year())}

	view.Render(w, "error", error)
}
