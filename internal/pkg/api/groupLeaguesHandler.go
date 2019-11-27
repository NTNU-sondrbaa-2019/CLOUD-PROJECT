package api

import (
    "encoding/json"
    "github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/HTTPErrors"
    "github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
    "log"
    "net/http"
    "strconv"
    "strings"
)

func GroupLeaguesHandler(w http.ResponseWriter, r *http.Request) HTTPErrors.Error {
    log.Println("Getting all leagues this group is a part of")
    urlPart := strings.Split(r.URL.Path, "/")

    groupID := urlPart[4]
    groupIDAsInt64, err := strconv.ParseInt(groupID, 10, 64)
    if err != nil {
        return HTTPErrors.NewError("Could not convert input to int", http.StatusBadRequest)
    }

    currentGroup, err := database.SelectGroupByID(groupIDAsInt64)
    if err != nil {
        return HTTPErrors.NewError("Could not get group by that ID from database", http.StatusInternalServerError)
    }
    league, err := database.SelectLeague(currentGroup.LeagueID)
    if err != nil {
        return HTTPErrors.NewError("Could not get league from group in database", http.StatusInternalServerError)
    }

    enc, err := json.Marshal(league)
    if err != nil {
        return HTTPErrors.NewError("Could not encode league struct to json", http.StatusInternalServerError)
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    _, err = w.Write(enc)
    if err != nil {
        return HTTPErrors.NewError("Could not write content to httpwriter", http.StatusInternalServerError)
    }

    return HTTPErrors.NewError("", 0)
}
