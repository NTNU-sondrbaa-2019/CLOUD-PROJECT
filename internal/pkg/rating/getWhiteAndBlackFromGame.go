package rating

// Functions loops through all team members to find the `teamMember` object of the `black` and `white` members of a game match.
func getWhiteAndBlackFromGame(game Game, teamMembers [] TeamMember) []TeamMember {
	var ret [] TeamMember
	var white TeamMember
	var black TeamMember
	for i := 0; i < len(teamMembers); i++ {
		if game.Players.White.User.Name == teamMembers[i].Username {
			white = teamMembers[i]
		} else if game.Players.Black.User.Name == teamMembers[i].Username {
			black = teamMembers[i]
		}
	}
	ret = append(ret, white)
	ret = append(ret, black)
	return ret
}
