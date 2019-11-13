package gauth

import (
    "context"
    "crypto/rand"
    "encoding/base64"
    "encoding/json"
    "github.com/pkg/errors"
    "log"
    "net/http"
    "time"
)

func generateStateOauthCookie(w http.ResponseWriter) string{
    // Set expiration of the cookie to 24 hours ahead of now
    var expiration = time.Now().Add(365 * 24 *   time.Hour)

    b := make([]byte, 16)
    rand.Read(b)
    state := base64.URLEncoding.EncodeToString(b)
    authCookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}

    http.SetCookie(w, &authCookie)

    // Return state so we can compare if we get the same back from gauth later
    return state
}

func getUserDataFromGoogle(code string) (userInfo, error) {
    var tempUser userInfo

    token, err := googleOauthConfig.Exchange(context.Background(), code)
    if err != nil {
        log.Println("Code exchange failed: " + err.Error())
        return tempUser, errors.New("Code exchange failed")
    }
    response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
    if err != nil {
        log.Println("Failed getting user info: " + err.Error())
        return tempUser, errors.New("Failed getting user info")
    }
    defer response.Body.Close()

    err = json.NewDecoder(response.Body).Decode(&tempUser)
    if err != nil {
        log.Println("Failed decoding user info from google: " + err.Error())
        return tempUser, errors.New("Failed decoding user info from google")
    }

    return tempUser, nil
}