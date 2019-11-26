package rating

import (
	"github.com/NTNU-sondrbaa-2019/CLOUD-PROJECT/internal/pkg/database"
	"log"
	"math"
	"strconv"
	"time"
)

func calculateElo(games []Game, teamMembers []TeamMember) []TeamMember {

	newTeamMembers := teamMembers

	var results [] database.RESULT
	var lichessIDS [] string
	// Loops through all games and calculates elo
	for i := 0; i < len(games); i++ {

		tmp := getWhiteAndBlackFromGame(games[i], newTeamMembers)
		white := tmp[0]
		black := tmp[1]
		group1, _ := database.SelectGroupByLeagueIDAndLeagueSeasonNameAndName(LICHESS_LEAGUE_ID, LICHESS_LEAGUE_SEASON_NAME, white.Username)
		group2, _ := database.SelectGroupByLeagueIDAndLeagueSeasonNameAndName(LICHESS_LEAGUE_ID, LICHESS_LEAGUE_SEASON_NAME, black.Username)

		//TODO get elo from own database
		count1, err1 := database.SelectResultCountByGroupID(group1.ID)
		count2, err2 := database.SelectResultCountByGroupID(group2.ID)
		log.Println("count 1 then 2")
		log.Println(count1, count2)

		if err1 != nil {
			log.Println(err1)
		}
		if err2 != nil {
			log.Println(err2)
		}

		if count1 != nil && *count1 > 0 {
			lastResultWhite, err := database.SelectResultLastByGroupId(group1.ID)
			if err != nil {
				log.Println(err)
			}
			if lastResultWhite != nil && err == nil {
				log.Println("Elo of resultwhite: ")
				log.Println(lastResultWhite.ELOAfter)
				log.Println("Result found, setting elo for this user to " + strconv.Itoa(lastResultWhite.ELOAfter))
				white.Elo = float64(lastResultWhite.ELOAfter)
			}
		} else {
			white.Elo = 1500
			log.Println("No results for this user, elo set to 1500")
		}

		if count2 != nil && *count2 > 0{
			lastResultBlack, err := database.SelectResultLastByGroupId(group1.ID)
			if err != nil {
				log.Println(err)
			}
			if lastResultBlack != nil && err == nil {
				log.Println("Elo of resultblack: ")
				log.Println(lastResultBlack.ELOAfter)
				log.Println("Result found, setting elo for this user to " + strconv.Itoa(lastResultBlack.ELOAfter))
				black.Elo = float64(lastResultBlack.ELOAfter)

			}
		} else {
			black.Elo = 1500
			log.Println("No results for this user, elo set to 1500")
		}

		// TODO GET ELO from last results corresponding to the teamMember/GROUP
		var w float64
		var b float64

		// TODO START make this block a single function to calculate elo
		K := 64.0

		R1 := math.Pow(10, white.Elo/400)
		R2 := math.Pow(10, black.Elo/400)

		WhiteChance := R1 / (R1 + R2)
		BlackChance := R2 / (R1 + R2)

		var outcome1 string
		var outcome2 string
		if games[i].Winner == "white" {
			w = white.Elo + (K * (1 - WhiteChance))
			b = black.Elo + (K * (0 - BlackChance))
			outcome1 = "win"
			outcome2 = "loss"
		} else if games[i].Winner == "black" {
			w = white.Elo + (K * (0 - WhiteChance))
			b = black.Elo + (K * (1 - BlackChance))
			outcome1 = "loss"
			outcome2 = "win"
		} else if games[i].Status == "draw"{
			w = white.Elo + (K * (0.5 - WhiteChance))
			b = black.Elo + (K * (0.5 - BlackChance))
			outcome1 = "draw"
			outcome2 = "draw"
		}

		// TODO END

		newTeamMembers = insertElo(w, b, white, black, newTeamMembers)

		var result1 database.RESULT
		var result2 database.RESULT

		result1.GroupID = group1.ID // DATABASE GROUP ID WTF DUDE
		result2.GroupID = group2.ID  // DATABASE GROUP ID WTF DUDE
		result1.ELOBefore = int(white.Elo)
		result2.ELOBefore = int(black.Elo)
		result1.ELOAfter = int(w)
		result2.ELOAfter = int(b)
		log.Println(result1.ELOBefore, result1.ELOAfter)
		log.Println(result2.ELOBefore, result2.ELOAfter)
		result1.Outcome = outcome1 // win lose draw
		result2.Outcome = outcome2 // win lose draw
		seconds := games[i].CreatedAt / 1000 		// Seconds
		nanoseconds := games[i].CreatedAt % 1000000	// Nanoseconds
		result1.Played = time.Unix(int64(seconds), int64(nanoseconds)) // time, unix convert
		result2.Played = result1.Played // time, unix convert

		results = append(results, result1)
		results = append(results, result2)
		lichessIDS = append(lichessIDS, games[i].ID)
		lichessIDS = append(lichessIDS, games[i].ID)
	}

	for i := 0; i < len(results); i++ {
		// TODO see if we get the result id from this function. Add result to RESULT_PLATFORM_ELO
		log.Println("Insert before: " + strconv.Itoa(results[i].ELOBefore) + " after: " + strconv.Itoa(results[i].ELOAfter) + " user: " + strconv.Itoa(int(results[i].GroupID)))
		id, _ := database.InsertResult(results[i])
		var resultPlatformElo database.RESULT_PLATFORM_ELO
		resultPlatformElo.PlatformEloID = LICHESS_PLATFORM_ID
		if id != nil{
			resultPlatformElo.ResultID = *id
		}
		resultPlatformElo.VerificationKey = lichessIDS[i]
		_, _ = database.InsertResultPlatformElo(resultPlatformElo)
	}

	for i := 0; i < len(newTeamMembers); i++ {
		print("Elo for member " + newTeamMembers[i].Username + ":\t" + strconv.FormatFloat(newTeamMembers[i].Elo, 'f', -1, 64) + "\n")
	}

	return newTeamMembers
}



