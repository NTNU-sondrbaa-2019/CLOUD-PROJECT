package server

import (
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/api"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/gauth"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/handler"
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/index"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler.MakeHandler(index.HandleIndex))
	http.HandleFunc("/login/", handler.MakeHandler(gauth.LoginHandler))
	http.HandleFunc("/loggedin/", handler.MakeHandler(gauth.LoggedInHandler))
	http.HandleFunc("/logout/", handler.MakeHandler(gauth.LogoutHandler))
	http.HandleFunc("/oauth2callback/", handler.MakeHandler(gauth.OauthCallBackHandler))
	//http.HandleFunc("/api/v1/", handler.MakeHandler(api.HandleAPI))
	http.HandleFunc("/api/v1/group/", handler.MakeHandler(api.GroupHandler))
	http.HandleFunc("/api/v1/user/", handler.MakeHandler(api.UserHandler))
	http.HandleFunc("/api/v1/league/", handler.MakeHandler(api.LeagueHandler))
	http.HandleFunc("/season-table/", handler.MakeHandler(index.EloDisplay))
	//http.HandleFunc("/api/v1/diagnostics/", handler.MakeHandler(api.DiagnosticsHandler))
}
