package gauth

import (
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "log"
    "net/http"
    "os"
)

var googleOauthConfig = &oauth2.Config{
    RedirectURL:    os.Getenv("POST_AUTH_REROUTE_URL"),
    ClientID:       os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
    ClientSecret:   os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
    Scopes:         []string{"https://www.googleapis.com/auth/userinfo.email"},
    Endpoint:       google.Endpoint,
}

func LoginHandler(w http.ResponseWriter, r *http.Request, title string) {
    // We are only interested in the error, because it can only be nil if the cookie exists or http.ErrNoCookie if the
    // cookie does not exist
    _, err := r.Cookie("sessionID")
    // If the sessionID cookie exists, redirect to the logged in page
    if err != http.ErrNoCookie {
        http.Redirect(w, r, "/loggedin", http.StatusTemporaryRedirect)
        return
    }
    // If there is no sessionID cookie, go through authentication
    log.Print("No sessionID cookie here, going straight to authentication")

    // Create oauthState cookie
    oauthState := generateStateOauthCookie(w)
    authURL := googleOauthConfig.AuthCodeURL(oauthState)
    http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}