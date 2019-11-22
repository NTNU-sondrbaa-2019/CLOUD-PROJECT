package database

import (
	"time"
)

type USER struct {
	ID         int64     `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Email      string    `json:"email" db:"email"`
	Registered time.Time `json:"registered" db:"registered"`
	LastOnline time.Time `json:"last_online" db:"last_online"`
}

func SelectUser(id int64) (*USER, error) {
	sth, err := connection.Prepare("SELECT * FROM USER WHERE id = ?")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	var user *USER

	err = sth.QueryRow(id).Scan(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func SelectUsers(where string) (*[]USER, error) {
	rows, err := connection.Query("SELECT * FROM USER " + where)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var user *USER
		err = rows.Scan(&user)

		if err != nil {
			return nil, err
		}

	}

	return nil, rows.Err()
}

func InsertUser(user USER) (*int64, error) {

	sth, err := connection.Prepare("INSERT INTO USER VALUES (NULL, ?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	result, err := sth.Exec(user.Name, user.Email, user.Registered, user.LastOnline)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func ModifyUser(id int64, user USER) error {

	sth, err := connection.Prepare("UPDATE USER SET name = ?, email = ?, registered = ?, last_online = ? WHERE id = ?")

	if err != nil {
		return err
	}

	defer sth.Close()

	_, err = sth.Exec(user.Name, user.Email, user.Registered, user.LastOnline, id)

	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(id int64) error {

	sth, err := connection.Prepare("DELETE FROM USER WHERE id = ?")

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
