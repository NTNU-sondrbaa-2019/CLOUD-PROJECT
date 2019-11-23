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

func SelectUserByID(id int64) (*USER, error) {
	sth, err := connection.Preparex("SELECT * FROM USER WHERE id = ?")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	var user USER
	err = sth.QueryRowx(id).StructScan(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func SelectUserByEmail(email string) (*USER, error) {
	sth, err := connection.Preparex("SELECT * FROM USER WHERE email = ?")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	var user USER
	err = sth.QueryRowx(email).StructScan(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func SelectUsers(where string) (*[]USER, error) {

	var users []USER

	rows, err := connection.Queryx("SELECT * FROM USER " + where)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var user USER
		err = rows.StructScan(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, user)

	}

	return &users, rows.Err()
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

	if len(user.Name) > 0 {

		sth, err := connection.Prepare("UPDATE USER SET name = ? WHERE id = ?")

		if err != nil {
			return err
		}

		defer sth.Close()

		_, err = sth.Exec(user.Name, id)

		if err != nil {
			return err
		}

	}

	if len(user.Email) > 0 {

		sth, err := connection.Prepare("UPDATE USER SET email = ? WHERE id = ?")

		if err != nil {
			return err
		}

		defer sth.Close()

		_, err = sth.Exec(user.Email, id)

		if err != nil {
			return err
		}

	}

	if !user.LastOnline.IsZero() {

		sth, err := connection.Prepare("UPDATE USER SET last_online = ? WHERE id = ?")

		if err != nil {
			return err
		}

		defer sth.Close()

		_, err = sth.Exec(user.LastOnline, id)

		if err != nil {
			return err
		}

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
