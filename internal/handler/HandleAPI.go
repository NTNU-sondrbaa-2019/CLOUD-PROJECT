package handler

import (
    "fmt"
    "github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/gauth"
    "net/http"
)

// HandleAPI - All requests to API endpoints will be handled here, follow testHandler for example
func HandleAPI(w http.ResponseWriter, r *http.Request, title string) {
    //http.HandleFunc("/api/v1/test/", MakeHandler(testHandler)) // Should be removed in final version
    http.HandleFunc("/api/v1/login/", MakeHandler(gauth.LoginHandler))
    http.HandleFunc("/api/v1/loggedin/", MakeHandler(gauth.LoggedInHandler))
    http.HandleFunc("/api/v1/logout/", MakeHandler(gauth.LogoutHandler))

    http.HandleFunc("/api/v1/authcallback/", MakeHandler(gauth.OauthCallBackHandler))

    // Here we could make a html documentation file for our API for now it returns not found.
    http.NotFound(w, r)
}

// TestHandler func can also be removed in final version . Currently here for you to know how to use
func testHandler(w http.ResponseWriter, r *http.Request, title string) {
    fmt.Println("Test Handler was called")
}