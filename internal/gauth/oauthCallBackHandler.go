package gauth

import (
    "log"
    "net/http"
)

func OauthCallBackHandler(w http.ResponseWriter, r *http.Request) {
    // Read state from cookie
    oauthState, _ := r.Cookie("oauthstate")

    // Compare state of callback to our local state
    if r.FormValue("state") != oauthState.Value {
        log.Print("Invalid google oauth state")
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }

    // Get our user's data from google
    tempUser, err := getUserDataFromGoogle(r.FormValue("code"))
    if err != nil {
        log.Println(err.Error())
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }

    // Save our user's info to the struct in memory
    dbSave(tempUser)

    // Make a cookie with our user's email lasting 1 hour
    userCookie := http.Cookie{
        Name:       "email",
        Value:      tempUser.Email,
        MaxAge:     3600,
    }
    http.SetCookie(w, &userCookie)

    // Print all users we have registered
    dbPrintAll(w)
}