package rating

// Inserts the new elo of white and black player into `teamMembers` and returns it.
func insertElo(eloWhite float64, eloBlack float64, white TeamMember, black TeamMember, teamMembers [] TeamMember) []TeamMember {
	for i := 0; i < len(teamMembers); i++ {
		if white.Username == teamMembers[i].Username {
			teamMembers[i].Elo = eloWhite
		} else if black.Username == teamMembers[i].Username {
			teamMembers[i].Elo = eloBlack
		}
	}
	return teamMembers
}
