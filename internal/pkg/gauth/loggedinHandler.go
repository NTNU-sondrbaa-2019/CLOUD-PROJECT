package gauth

import (
    "fmt"
    "github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/HTTPErrors"
    "github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
    "net/http"
)

func LoggedInHandler(w http.ResponseWriter, r *http.Request) HTTPErrors.Error {
    fmt.Fprintln(w, "Here you see your email and Lichess key. You are redirected here automatically if you are already logged in.")
    sessionID := GetCookieValueByName(r.Cookies(),"sessionID")
    fmt.Fprintln(w, "You are logged in with session ID: " + sessionID)
    fmt.Fprintln(w, "Your information:")

    dbPrintSpecificID(w, sessionID)

    allUsers, err := database.SelectUsers("")
    if err != nil {
        return HTTPErrors.NewError("Could not select all users from database", http.StatusInternalServerError)
    }

    fmt.Fprintln(w, "Every user:")
    fmt.Fprint(w, allUsers)
    return HTTPErrors.NewError("", 0)
}