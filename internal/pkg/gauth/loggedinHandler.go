package gauth

import (
    "fmt"
    "github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/HTTPErrors"
    "github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
    "net/http"
)

func LoggedInHandler(w http.ResponseWriter, r *http.Request) HTTPErrors.Error {
    fmt.Fprintln(w, "Here you see your email name.")
    sessionID := GetCookieValueByName(r.Cookies(),"sessionID")
    fmt.Fprintln(w, "You are logged in with session ID: " + sessionID)
    fmt.Fprintln(w, "Your information:")

    if ExistsUserSession(sessionID) {
        userID, err := GetUserIDFromSessionID(sessionID)
        if err != nil {
            return HTTPErrors.NewError("Failed getting userID from database using sessionID",
                http.StatusInternalServerError)
        }
        userInfo, err := database.SelectUserByID(*userID)
        fmt.Fprintln(w, userInfo)
    }

    fmt.Fprintln(w, "Every user:")
    allUsers, err := database.SelectUsers("")
    if err != nil {
        return HTTPErrors.NewError("Could not select all users from database", http.StatusInternalServerError)
    }
    fmt.Fprint(w, allUsers)

    return HTTPErrors.NewError("", 0)
}