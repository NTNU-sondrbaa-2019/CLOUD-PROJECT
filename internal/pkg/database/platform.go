package database

type PLATFORM struct {
	ID      int64  `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Version string `json:"version" db:"version"`
}

func SelectPlatform(id int64) (*PLATFORM, error) {
	sth, err := connection.Preparex("SELECT * FROM PLATFORM WHERE id = ?")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	var platform PLATFORM
	err = sth.QueryRowx(id).StructScan(&platform)

	if err != nil {
		return nil, err
	}

	return &platform, nil
}

func SelectPlatformIDByNameAndVersion(name string, version string) (*int64, error) {
	sth, err := connection.Prepare("SELECT id FROM PLATFORM WHERE name = ? AND version = ?")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	var id int64
	err = sth.QueryRow(name, version).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func SelectPlatforms(where string) (*[]PLATFORM, error) {

	var platforms []PLATFORM

	rows, err := connection.Queryx("SELECT * FROM PLATFORM " + where)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var platform PLATFORM
		err = rows.StructScan(&platform)

		if err != nil {
			return nil, err
		}

		platforms = append(platforms, platform)

	}

	return &platforms, rows.Err()
}

func InsertPlatform(platform PLATFORM) (*int64, error) {

	sth, err := connection.Prepare("INSERT INTO PLATFORM VALUES (NULL, ?, ?)")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	result, err := sth.Exec(platform.Name, platform.Version)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func DeletePlatform(id int64) error {

	sth, err := connection.Prepare("DELETE FROM PLATFORM WHERE id = ?")

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
