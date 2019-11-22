package handler

import (
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/HTTPErrors"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/view"
	"net/http"
)


func MakeHandler(fn func(http.ResponseWriter, *http.Request) HTTPErrors.Error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err.Message != "" {
			view.ErrorPage(w, err.Message, err.Code)
		}
	}
}
