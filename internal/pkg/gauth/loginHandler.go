package gauth

 import (
     "github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/HTTPErrors"
     "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "net/http"
    "os"
)

var GoogleOauthConfig = &oauth2.Config{
    RedirectURL:    os.Getenv("POST-AUTH-REDIRECT-URL"),
    ClientID:       os.Getenv("GOOGLE-OAUTH-CLIENT-ID"),
    ClientSecret:   os.Getenv("GOOGLE-OAUTH-CLIENT-SECRET"),
    Scopes:         []string{"profile"},
    Endpoint:       google.Endpoint,
}

func LoginHandler(w http.ResponseWriter, r *http.Request) HTTPErrors.Error {
    // Create oauthState cookie
    oauthState := generateStateOauthCookie(w)
    authURL := GoogleOauthConfig.AuthCodeURL(oauthState)
    http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)

    return HTTPErrors.NewError("", 0)
}