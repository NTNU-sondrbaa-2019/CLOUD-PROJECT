# `rating` package

## Public

### `func GetTeamElo(teamIdKey string)`
Only function that should be called from this package.

Parameter `teamIdKey` is the id of a team on lichess.org

Does a GET request to lichess and gets all the members in a team. Calls `GetGamesInTeam` to get all games in a team. 
Calls `sortGames` to sort these games. Calls `calculateElo` to calculate Elo. 

returns `[]TeamMember` with all team members in a team with an internal ELO in team.

## Private

### `func getGamesInTeam(teamMembers [] TeamMember)`
Parameter `teamMembers` is a slice of team members

Loops through every team member vs team member permutation and calls `getGamesOfMemberVSMember`.

returns `[]Game` of every match between teammembers in a team.

### `func getGamesOfMemberVSMember(member TeamMember, vsMember TeamMember)`
Parameters `member` and `vsMember`, two members which we get all games between.

returns `[]Game` which contains all matches between two members.

### `func getRequest(c *http.Client, s string)`
Parameters `c` and `s`, where `s` is the GET request address.

Is a simple function for get requests.

returns `*http.Response`

### `func getWhiteAndBlackFromGame(game Game, teamMembers [] TeamMember)`
Parameters `game` is a game object, `teamMembers` is a slice of all team members.

Functions loops through all team members to find the `teamMember` object of the `black` and `white` members of a game match.

returns `[]TeamMember` which contains White and Black corresponding teamMember objects. White in the first index and black in the second.

### `func calculateElo(games []Game, teamMembers []TeamMember)`
Parameters `games` and `teamMembers`. `games` is a slice of games and `teamMembers` is a slice of all team members in a team.

Calculates the elo of team members match by match (game). 

returns `[]TeamMember` with all elo scores updated

### `func insertElo(eloWhite float64, eloBlack float64, white TeamMember, black TeamMember, teamMembers [] TeamMember)`
Parameters  `eloWhite`, `eloBlack`, `white`, `black`, `teamMembers`

Inserts the new elo of white and black player into `teamMembers` and returns it.

returns `[]TeamMember` with updated elo.

### `func sortGames(games []Game)`
Parameter `games`

Sorts games by timestamp and returns

returns `[]Game`

### `func fakeTeamMembers`
Returns fake test data `[]TeamMembers`

### `structs`

    type Game struct {
    	ID        string `json:"id"`
    	CreatedAt int    `json:"createdAt"`
    	Players   struct {
    		White struct {
    			User struct {
    				ID   string `json:"id"`
    				Name string `json:"name"`
    			} `json:"user"`
    		} `json:"white"`
    		Black struct {
    			User struct {
    				ID   string `json:"id"`
    				Name string `json:"name"`
    			} `json:"user"`
    		} `json:"black"`
    	} `json:"players"`
    	Winner string `json:"winner"`
    }


    type TeamMember struct {
    	ID                string  `json:"id"`
    	Username          string  `json:"username"`
    	InternalCreatedAt int     `json:"internalCreatedAt"`
    	Elo               float64 `json:"elo"`
    }
