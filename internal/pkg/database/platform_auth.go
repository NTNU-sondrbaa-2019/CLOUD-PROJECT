package database

type PLATFORM_AUTH struct {
	PlatformID int64 `json:"platform_id" db:"platform_id"`
}

func PlatformAuthExists(platform_id int64) (*bool, error) {

	sth, err := connection.Prepare("SELECT COUNT(*) FROM PLATFORM_AUTH WHERE platform_id = ?")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	var count int
	err = sth.QueryRow(platform_id).Scan(&count)

	if err != nil {
		return nil, err
	}

	flag := count > 0

	return &flag, nil
}

func InsertPlatformAuth(platform_auth PLATFORM_AUTH) (*int64, error) {

	sth, err := connection.Prepare("INSERT INTO PLATFORM_AUTH VALUES (?)")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	result, err := sth.Exec(platform_auth.PlatformID)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func DeletePlatformAuth(platform_id int64) error {

	sth, err := connection.Prepare("DELETE FROM PLATFORM_AUTH WHERE platform_id = ?")

	if err != nil {
		return err
	}

	defer sth.Close()

	_, err = sth.Exec(platform_id)

	if err != nil {
		return err
	}

	return nil
}
