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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    // Check if the sessionID cookie exists
    sessionIDCookie, err := r.Cookie("sessionID")
    if err != nil {
        log.Print("No sessionID cookie here, going straight to authentication")
    } else if sessionIDCookie.Value != "" {
        // If the sessionID is found in the database, redirect to logged in page
        if dbFoundSession(sessionIDCookie.Value){
            http.Redirect(w, r, "/loggedin", http.StatusPermanentRedirect)
        }
    }

    // If there is no userID cookie, go through authentication
    // Create oauthState cookie
    oauthState := generateStateOauthCookie(w)
    authURL := googleOauthConfig.AuthCodeURL(oauthState)
    http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}