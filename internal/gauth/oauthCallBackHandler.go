package gauth

import (
    "crypto/rand"
    "encoding/base64"
    "log"
    "net/http"
    "time"
)

type userInfoFromGoogle struct {
    Email       string      `json:"email"`
}

func OauthCallBackHandler(w http.ResponseWriter, r *http.Request, title string) {
    log.Println("------------------------------In OauthCallbackHandler")

    // Read state from cookie
    // oauthState, _ := r.Cookie("oauthstate")

    /* Compare state of callback to our local state
    if r.FormValue("state") != oauthState.Value {
        log.Print("Invalid google oauth state")
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }
    */

    // Get our user's data from google
    tempUserFromGoogle, err := getUserDataFromGoogle(r.FormValue("code"))
    if err != nil {
        log.Println(err.Error())
        http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
        return
    }

    //Make a random 16 characters long ID for this user
    b := make([]byte, 16)
    rand.Read(b)
    tempID := base64.URLEncoding.EncodeToString(b)

    // Make a cookie with our user's id that expires in 24 hours
    sessionIDCookie := http.Cookie{
        Name:       "sessionID",
        Value:      tempID,
        Expires:    time.Now().Add(24 * time.Hour),
    }
    r.AddCookie(&sessionIDCookie)

    // Make a tempUser with the info we got from google and our new sessionID
    tempUser := userInfo{
        Email:      tempUserFromGoogle.Email,
        LichessKey: "",
        LastSessionID: tempID,
    }

    // Save our user's info to the struct in memory
    dbSave(tempUser)

    // Now that the user is logged in, redirect to the logged in page
    http.Redirect(w, r, "/api/v1/loggedin/", http.StatusPermanentRedirect)
}