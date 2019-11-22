package gauth

import (
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "net/http"
    "os"
)

var GoogleOauthConfig = &oauth2.Config{
    RedirectURL:    os.Getenv("POST_AUTH_REROUTE_URL"),
    ClientID:       os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
    ClientSecret:   os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
    Scopes:         []string{"https://www.googleapis.com/auth/userinfo.email"},
    Endpoint:       google.Endpoint,
}

func LoginHandler(w http.ResponseWriter, r *http.Request, title string) {
    // Create oauthState cookie
    oauthState := generateStateOauthCookie(w)
    authURL := GoogleOauthConfig.AuthCodeURL(oauthState)
    http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}