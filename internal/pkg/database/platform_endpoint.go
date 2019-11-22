package database

type PLATFORM_ENDPOINT struct {
	PlatformID int64  `json:"platform_id" db:"platform_id"`
	Name       string `json:"name" db:"name"`
	Path       string `json:"path" db:"path"`
}

func SelectPlatformEndpointPath(platform_id int64, name string) (*string, error) {

	sth, err := connection.Prepare("SELECT path FROM PLATFORM_ENDPOINT WHERE platform_id = ? AND name = ?")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	var path string
	err = sth.QueryRow(platform_id, name).Scan(&path)

	if err != nil {
		return nil, err
	}

	return &path, nil
}

func SelectPlatformEndpoints(where string) (*[]PLATFORM_ENDPOINT, error) {

	var platform_endpoints []PLATFORM_ENDPOINT

	rows, err := connection.Queryx("SELECT * FROM PLATFORM_ENDPOINT " + where)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var platform_endpoint PLATFORM_ENDPOINT
		err = rows.StructScan(&platform_endpoint)

		if err != nil {
			return nil, err
		}

		platform_endpoints = append(platform_endpoints, platform_endpoint)

	}

	return &platform_endpoints, rows.Err()
}

func InsertPlatformEndpoint(platform_endpoint PLATFORM_ENDPOINT) (*int64, error) {

	sth, err := connection.Prepare("INSERT INTO PLATFORM_ENDPOINT VALUES (?, ?, ?)")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	result, err := sth.Exec(platform_endpoint.PlatformID, platform_endpoint.Name, platform_endpoint.Path)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func ModifyPlatformEndpointPath(platform_id int64, name string, path string) error {

	sth, err := connection.Prepare("UPDATE PLATFORM_ENDPOINT SET path = ? WHERE platform_id = ? AND name = ?")

	if err != nil {
		return err
	}

	defer sth.Close()

	_, err = sth.Exec(path, platform_id, name)

	if err != nil {
		return err
	}

	return nil
}

func DeletePlatformEndpoint(platform_id int64, name string) error {

	sth, err := connection.Prepare("DELETE FROM PLATFORM_ENDPOINT WHERE platform_id = ? AND name = ?")

	if err != nil {
		return err
	}

	defer sth.Close()

	_, err = sth.Exec(platform_id, name)

	if err != nil {
		return err
	}

	return nil
}
