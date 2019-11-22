package view

import "html/template"

var Templates = template.Must(template.ParseFiles(
	"web/static/login.html",
	"web/static/ucp.html",
	"web/static/error.html",
	))
