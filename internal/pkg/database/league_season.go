package database

import "time"

type LEAGUE_SEASON struct {
	LeagueID int64     `json:"league_id" db:"league_id"`
	Name     string    `json:"name" db:"name"`
	Start    time.Time `json:"start" db:"start"`
	End      time.Time `json:"end" db:"end"`
}

func SelectLeagueSeason(league_id int64, name string) (*LEAGUE_SEASON, error) {
	sth, err := connection.Preparex("SELECT * FROM LEAGUE_SEASON WHERE league_id = ? AND name = ?")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	var league_season LEAGUE_SEASON
	err = sth.QueryRowx(league_id, name).StructScan(&league_season)

	if err != nil {
		return nil, err
	}

	return &league_season, nil
}

func InsertLeagueSeason(league_season LEAGUE_SEASON) (*int64, error) {

	sth, err := connection.Prepare("INSERT INTO LEAGUE_SEASON VALUES (?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	result, err := sth.Exec(league_season.LeagueID, league_season.Name, league_season.Start, league_season.End)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func ModifyLeagueSeason(league_id int64, name string, league_season LEAGUE_SEASON) error {

	if !league_season.Start.IsZero() {

		sth, err := connection.Prepare("UPDATE LEAGUE_SEASON SET start = ? WHERE league_id = ? AND name = ?")

		if err != nil {
			return err
		}

		defer sth.Close()

		_, err = sth.Exec(league_season.Start, league_id, name)

		if err != nil {
			return err
		}

	}

	if !league_season.End.IsZero() {

		sth, err := connection.Prepare("UPDATE LEAGUE_SEASON SET end = ? WHERE league_id = ? AND name = ?")

		if err != nil {
			return err
		}

		defer sth.Close()

		_, err = sth.Exec(league_season.End, league_id, name)

		if err != nil {
			return err
		}

	}

	return nil
}

func DeleteLeagueSeason(league_id int64, name string) error {

	sth, err := connection.Prepare("DELETE FROM LEAGUE_SEASON WHERE league_id = ? AND name = ?")

	if err != nil {
		return err
	}

	defer sth.Close()

	_, err = sth.Exec(league_id, name)

	if err != nil {
		return err
	}

	return nil
}
