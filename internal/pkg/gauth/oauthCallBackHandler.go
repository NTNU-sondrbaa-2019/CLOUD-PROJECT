package gauth

import (
    "crypto/rand"
    "encoding/base64"
    "github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/HTTPErrors"
    "github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
    "net/http"
    "time"
)

type userInfoFromGoogle struct {
    Email       string      `json:"email"`
    Name        string      `json:"name"`
}


func OauthCallBackHandler(w http.ResponseWriter, r *http.Request) HTTPErrors.Error{

    // Read state from cookie
    oauthState, _ := r.Cookie("oauthstate")

    // Compare state of callback to our local state
    if r.FormValue("state") != oauthState.Value {
        return HTTPErrors.NewError("Invalid state from google", http.StatusInternalServerError)
    }

    // Get our user's data from google
    tempUserFromGoogle, err := getUserDataFromGoogle(r.FormValue("code"))
    if err != nil {
        return HTTPErrors.NewError("Could not get user data from google", http.StatusInternalServerError)
    }

    //Make a random 16 characters long ID for this user
    b := make([]byte, 16)
    rand.Read(b)
    tempID := base64.URLEncoding.EncodeToString(b)

    // Make a cookie with our user's id that expires in 24 hours
    sessionIDCookie := http.Cookie{
        Name:       "sessionID",
        Path:       "/",
        Value:      tempID,
        Expires:    time.Now().Add(24 * time.Hour),
    }
    http.SetCookie(w, &sessionIDCookie)

    tempUser := database.USER{
        Name:       tempUserFromGoogle.Name,
        Email:      tempUserFromGoogle.Email,
        Registered: time.Now(),
        LastOnline: time.Now(),
    }

    userID, err :=  database.InsertUser(tempUser)
    // If err is not nil, then the user with this email already exists
    // So we get that user and update the lastonline time to time.Now()
    if err != nil {
        modUser, err := database.SelectUserByEmail(tempUser.Email)
        if err != nil {
            return HTTPErrors.NewError("Could not select existing user from database", http.StatusInternalServerError)
        }

        err = database.ModifyUser(modUser.ID, tempUser)
        if err != nil {
            return HTTPErrors.NewError("Could not modify existing user in database", http.StatusInternalServerError)
        }

        tempCacheSession := userSession{
            SessionID: sessionIDCookie.Value,
            UserID:    modUser.ID,
        }
        AddUserSession(tempCacheSession)
    } else {
        tempCacheSession := userSession{
            SessionID: sessionIDCookie.Value,
            UserID:    *userID,
        }
        AddUserSession(tempCacheSession)
    }

    // Now that the user is logged in, redirect to the logged in page
    http.Redirect(w, r, "/loggedin/", http.StatusPermanentRedirect)

    return HTTPErrors.NewError("", 0)
}