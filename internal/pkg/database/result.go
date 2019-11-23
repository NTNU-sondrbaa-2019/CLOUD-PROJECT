package database

import (
	"time"
)

type RESULT struct {
	ID            int64     `json:"id" db:"id"`
	GroupID       int64     `json:"group_id" db:"group_id"`
	ELOBefore     int       `json:"elo_before" db:"elo_before"`
	ELOAfter      int       `json:"elo_after" db:"elo_after"`
	ELODifference int       `json:"elo_difference" db:"elo_difference"`
	Outcome       string    `json:"outcome" db:"outcome"`
	Played        time.Time `json:"played" db:"played"`
}

func SelectCountResultByGroupID(id int64) (int, error) {
	sth, err := connection.Preparex("SELECT COUNT(*) FROM RESULT WHERE id = ?")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	var result int
	err = sth.QueryRowx(id).StructScan(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func SelectResultLastByGroupId(id int64) (*RESULT, error) {
	sth, err := connection.Preparex("SELECT * FROM RESULT WHERE id = ? ORDER BY played DESC LIMIT 1")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	var result RESULT
	err = sth.QueryRowx(id).StructScan(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func SelectResult(id int64) (*RESULT, error) {
	sth, err := connection.Preparex("SELECT * FROM RESULT WHERE id = ?")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	var result RESULT
	err = sth.QueryRowx(id).StructScan(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func SelectResults(where string) (*[]RESULT, error) {

	var results []RESULT

	rows, err := connection.Queryx("SELECT * FROM RESULT " + where)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var result RESULT
		err = rows.StructScan(&result)

		if err != nil {
			return nil, err
		}

		results = append(results, result)

	}

	return &results, rows.Err()
}

func InsertResult(result RESULT) (*int64, error) {

	sth, err := connection.Prepare("INSERT INTO RESULT VALUES (NULL, ?, ?, ?, DEFAULT, ?, ?)")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	r, err := sth.Exec(result.GroupID, result.ELOBefore, result.ELOAfter, result.Outcome, result.Played)

	if err != nil {
		return nil, err
	}

	id, err := r.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func DeleteResult(id int64) error {

	sth, err := connection.Prepare("DELETE FROM RESULT WHERE id = ?")

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
