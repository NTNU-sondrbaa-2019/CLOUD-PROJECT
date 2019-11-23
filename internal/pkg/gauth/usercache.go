package gauth

import (
    "encoding/json"
    "github.com/NTNU-sondrbaa-2019/CLOUD-O1/pkg/CO1Cache"
    "log"
)

type userSession struct {
    SessionID   string  `json:"session"`
    UserID      int64  `json:"user"`
}
func AddUserSession(session userSession) {
    log.Println("Saving new user session pair to cache")
    CO1Cache.WriteJSON(session.SessionID, session)
}

func ExistsUserSession(sessionID string) bool {
    return CO1Cache.Verify(sessionID)
}

func GetUserIDFromSessionID (sessionID string) (int64, error){
    var tempUserSession userSession
    err := json.Unmarshal(CO1Cache.Read(sessionID), &tempUserSession)
    if err != nil {
        return nil, err
    }
    return tempUserSession.UserID, nil
}