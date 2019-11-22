package database

type PLATFORM_ELO struct {
	PlatformID int64 `json:"platform_id" db:"platform_id"`
}

func PlatformEloExists(platform_id int64) (*bool, error) {

	sth, err := connection.Prepare("SELECT COUNT(*) FROM PLATFORM_ELO WHERE platform_id = ?")

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

func InsertPlatformElo(platform_elo PLATFORM_ELO) (*int64, error) {

	sth, err := connection.Prepare("INSERT INTO PLATFORM_ELO VALUES (?)")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	result, err := sth.Exec(platform_elo.PlatformID)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func DeletePlatformElo(platform_id int64) error {

	sth, err := connection.Prepare("DELETE FROM PLATFORM_ELO WHERE platform_id = ?")

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
