package database

import (
	"time"
)

type GROUP struct {
	ID               int64     `json:"id" db:"id"`
	LeagueID         int64     `json:"league_id" db:"league_id"`
	LeagueSeasonName string    `json:"league_season_name" db:"league_season_name"`
	Name             string    `json:"name" db:"name"`
	Created          time.Time `json:"created" db:"created"`
}

func SelectGroup(id int64) (*GROUP, error) {
	sth, err := connection.Preparex("SELECT * FROM `GROUP` WHERE id = ?")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	var group GROUP
	err = sth.QueryRowx(id).StructScan(&group)

	if err != nil {
		return nil, err
	}

	return &group, nil
}

func SelectGroupsByLeagueIDAndLeagueSeasonName(league_id int64, league_season_name string) (*[]GROUP, error) {

	var groups []GROUP

	sth, err := connection.Preparex("SELECT * FROM GROUP WHERE league_id = ? AND league_season_name = ?")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	rows, err := sth.Queryx(league_id, league_season_name)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var group GROUP
		err = rows.StructScan(&group)

		if err != nil {
			return nil, err
		}

		groups = append(groups, group)

	}

	return &groups, nil
}

func SelectGroups(where string) (*[]GROUP, error) {

	var groups []GROUP

	rows, err := connection.Queryx("SELECT * FROM `GROUP` " + where)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var group GROUP
		err = rows.StructScan(&group)

		if err != nil {
			return nil, err
		}

		groups = append(groups, group)

	}

	return &groups, rows.Err()
}

func InsertGroup(group GROUP) (*int64, error) {

	sth, err := connection.Prepare("INSERT INTO `GROUP` VALUES (NULL, ?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	result, err := sth.Exec(group.LeagueID, group.Name, group.LeagueSeasonName, group.Name, group.Created)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func ModifyGroup(id int64, group GROUP) error {

	if len(group.Name) > 0 {

		sth, err := connection.Prepare("UPDATE GROUP SET name = ? WHERE id = ?")

		if err != nil {
			return err
		}

		defer sth.Close()

		_, err = sth.Exec(group.Name, id)

		if err != nil {
			return err
		}

	}

	return nil
}

func DeleteGroup(id int64) error {

	sth, err := connection.Prepare("DELETE FROM GROUP WHERE id = ?")

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
