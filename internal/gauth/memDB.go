package gauth

import (
    "fmt"
    "net/http"
)

type userInfo struct {
    Email           string      `json:"email"`
    LichessKey      string      `json:"lichesskey"`
    LastSessionID   string      `json:"lastsessionid"`
}

var users []userInfo

// Save the user info to memory
func dbSave(in userInfo) {
    // If this email already exists in memory, only update the LastSessionID
    for _, i := range users {
        if i.Email == in.Email {
            i.LastSessionID = in.LastSessionID
            return
        }
    }

    // Append this new user to memory
    users = append(users, in)
}

// If the compareID is found as the sessionID anywhere in the database, return true
func dbFoundSession(compareID string) bool{
    for _, i := range users {
        if i.LastSessionID == compareID {
            return true
        }
    }
    return false
}

func dbPrintSpecificID(w http.ResponseWriter, sessionID string) {
    for _, i := range users {
        if i.LastSessionID == sessionID {
            fmt.Fprintln(w, i)
        }
    }
}

func dbPrintAll(w http.ResponseWriter) {
    fmt.Fprintln(w, users)
}