package gauth

import (
    "github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/HTTPErrors"
    "net/http"
    "time"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) HTTPErrors.Error {
    // Make the sessionID cookie max age of 0, it will delete itself
    sessionIDCookie := http.Cookie{
        Name:       "sessionID",
        Path:       "/",
        Value:      "",
        Expires:    time.Now(),
        MaxAge:     0,
    }
    http.SetCookie(w, &sessionIDCookie)

    // Redirect to main screen
    http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

    return HTTPErrors.NewError("", 0)
}