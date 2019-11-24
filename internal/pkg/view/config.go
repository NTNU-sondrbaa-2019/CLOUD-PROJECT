package view

import (
	"html/template"
	"os"
)

var GOPATH = os.Getenv("GOPATH")

var Templates = template.Must(template.ParseFiles(
	GOPATH + "/src/CLOUD-PROJECT/web/static/login.html",
	GOPATH + "/src/CLOUD-PROJECT/web/static/ucp.html",
	GOPATH + "/src/CLOUD-PROJECT/web/static/error.html",
	))
