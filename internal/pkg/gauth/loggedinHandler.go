package gauth

import (
    "fmt"
    "net/http"
)

func LoggedInHandler(w http.ResponseWriter, r *http.Request, title string) {
    fmt.Fprintln(w, "Here you see your email and Lichess key. You are redirected here automatically if you are already logged in.")
    sessionID := GetCookieValueByName(r.Cookies(),"sessionID")
    fmt.Fprintln(w, "You are logged in with session ID: " + sessionID)
    fmt.Fprintln(w, "Your information:")

    dbPrintSpecificID(w, sessionID)

    fmt.Fprintln(w, "Every user:")
    dbPrintAll(w)
}