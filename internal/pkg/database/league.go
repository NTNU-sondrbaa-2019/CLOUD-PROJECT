package database

import (
	"time"
)

type LEAGUE struct {
	ID             int64     `json:"id" db:"id"`
	PlatformEloID  int64     `json:"platform_elo_id" db:"platform_elo_id"`
	Name           string    `json:"name" db:"name"`
	GroupsLimit    int       `json:"groups_limit" db:"groups_limit"`
	GroupSizeLimit int       `json:"group_size_limit" db:"group_size_limit"`
	Start          time.Time `json:"start" db:"start"`
	End            time.Time `json:"end" db:"end"`
}

func LeagueExists(leagueName string) (bool, error) {
	sth, err := connection.Prepare("SELECT COUNT(*) FROM LEAGUE WHERE name = '?'")

	if err != nil {
		return false, err
	}

	defer sth.Close()

	var count int
	err = sth.QueryRow(leagueName).Scan(&count)

	if err != nil {
		return false, err
	}

	flag := count > 0

	return flag, nil
}

func SelectLeague(id int64) (*LEAGUE, error) {
	sth, err := connection.Preparex("SELECT * FROM LEAGUE WHERE id = ?")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	var league LEAGUE
	err = sth.QueryRowx(id).StructScan(&league)

	if err != nil {
		return nil, err
	}

	return &league, nil
}

func SelectLeagueByPlatformEloIDAndName(platform_elo_id int64, name string) (*LEAGUE, error) {

	sth, err := connection.Preparex("SELECT * FROM LEAGUE WHERE platform_elo_id = ? AND name = ?")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	var league LEAGUE
	err = sth.QueryRowx(platform_elo_id, name).StructScan(&league)

	if err != nil {
		return nil, err
	}

	return &league, nil
}

func SelectLeagues(where string) (*[]LEAGUE, error) {

	var leagues []LEAGUE

	rows, err := connection.Queryx("SELECT * FROM LEAGUE " + where)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var league LEAGUE
		err = rows.StructScan(&league)

		if err != nil {
			return nil, err
		}

		leagues = append(leagues, league)

	}

	return &leagues, rows.Err()
}

func InsertLeague(league LEAGUE) (*int64, error) {

	sth, err := connection.Prepare("INSERT INTO LEAGUE VALUES (NULL, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	result, err := sth.Exec(league.PlatformEloID, league.Name, league.GroupsLimit, league.GroupSizeLimit, league.Start, league.End)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func ModifyLeague(id int64, league LEAGUE) error {

	if len(league.Name) > 0 {

		sth, err := connection.Prepare("UPDATE LEAGUE SET name = ? WHERE id = ?")

		if err != nil {
			return err
		}

		defer sth.Close()

		_, err = sth.Exec(league.Name, id)

		if err != nil {
			return err
		}

	}

	if league.GroupsLimit > 0 {

		sth, err := connection.Prepare("UPDATE LEAGUE SET groups_limit = ? WHERE id = ?")

		if err != nil {
			return err
		}

		defer sth.Close()

		_, err = sth.Exec(league.GroupsLimit, id)

		if err != nil {
			return err
		}

	}

	if league.GroupSizeLimit > 0 {

		sth, err := connection.Prepare("UPDATE LEAGUE SET group_size_limit = ? WHERE id = ?")

		if err != nil {
			return err
		}

		defer sth.Close()

		_, err = sth.Exec(league.GroupSizeLimit, id)

		if err != nil {
			return err
		}

	}

	if !league.Start.IsZero() {

		sth, err := connection.Prepare("UPDATE LEAGUE SET start = ? WHERE id = ?")

		if err != nil {
			return err
		}

		defer sth.Close()

		_, err = sth.Exec(league.Start, id)

		if err != nil {
			return err
		}

	}

	if !league.End.IsZero() {

		sth, err := connection.Prepare("UPDATE LEAGUE SET end = ? WHERE id = ?")

		if err != nil {
			return err
		}

		defer sth.Close()

		_, err = sth.Exec(league.End, id)

		if err != nil {
			return err
		}

	}

	return nil
}

func DeleteLeague(id int64) error {

	sth, err := connection.Prepare("DELETE FROM LEAGUE WHERE id = ?")

	if err != nil {
		return err
	}

	defer sth.Close()

	_, err = sth.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
