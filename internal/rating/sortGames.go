package rating

import (
	"sort"
	"strconv"
)

func SortGames(games []Game) []Game{
	sort.Slice(games, func(i, j int) bool { return games[i].CreatedAt < games[j].CreatedAt })
	// TODO remove print
	for i := 0; i < len(games); i++ {
		print("Length of game " + strconv.Itoa(i) + ": " + strconv.Itoa(games[i].CreatedAt) + "\n")
	}
	return games
}