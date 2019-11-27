package view

import (
	"html/template"
	"os"
)

var GOPATH = os.Getenv("GOPATH")

var Templates = template.Must(template.ParseFiles(
	"web/static/login.html",
	"web/static/ucp.html",
	"web/static/error.html",
	"web/static/season_table.html",
))
