package gauth

import (
    "fmt"
    "net/http"
)

func LoggedInHandler(w http.ResponseWriter, r *http.Request, title string) {
    fmt.Fprintln(w, "Here you see your email and Lichess key. You are redirected here automatically if you are already logged in.")
    sessionIDCookie, _ := r.Cookie("sessionID")
    fmt.Fprintln(w, "You are logged in with session ID: " + sessionIDCookie.Value)
    fmt.Fprintln(w, "Your information:")

    dbPrintSpecificID(w, sessionIDCookie.Value)

    fmt.Fprintln(w, "Every user:")
    dbPrintAll(w)
}