package gauth

import (
    "net/http"
    "time"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request, title string) {
    // Make the sessionID cookie max age of 0, it will delete itself
    sessionIDCookie := http.Cookie{
        Name:       "sessionID",
        Value:      "",
        Expires:    time.Now(),
        MaxAge:     0,
    }
    http.SetCookie(w, &sessionIDCookie)

    // Redirect to main screen
    http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}