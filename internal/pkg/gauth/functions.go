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
    // Set expiration of the cookie to a year from now
    var expiration = time.Now().Add(365 * 24 *   time.Hour)

    b := make([]byte, 16)
    rand.Read(b)
    state := base64.URLEncoding.EncodeToString(b)

    authCookie := http.Cookie{Name: "oauthstate", Path: "/", Value: state, Expires: expiration}

    http.SetCookie(w, &authCookie)

    // Return state so we can compare if we get the same back from gauth later
    return state
}

// Gets the logged in user's email adress we use for identification
func getUserDataFromGoogle(code string) (userInfoFromGoogle, error) {
    var tempUser userInfoFromGoogle

    token, err := GoogleOauthConfig.Exchange(context.Background(), code)
    if err != nil {
        log.Print("Code exchange failed: " + err.Error())
        return tempUser, errors.New("Code exchange failed")
    }
    response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
    if err != nil {
        log.Print("Failed getting user info: " + err.Error())
        return tempUser, errors.New("Failed getting user info")
    }
    defer response.Body.Close()

    err = json.NewDecoder(response.Body).Decode(&tempUser)
    if err != nil {
        log.Print("Failed decoding user info from google: " + err.Error())
        return tempUser, errors.New("Failed decoding user info from google")
    }

    return tempUser, nil
}

func GetCookieValueByName(cookie []*http.Cookie, name string) string {
    cookieLen := len(cookie)
    result := ""
    for i := 0; i < cookieLen; i++ {
        if cookie[i].Name == name {
            result = cookie[i].Value
        }
    }
    return result
}