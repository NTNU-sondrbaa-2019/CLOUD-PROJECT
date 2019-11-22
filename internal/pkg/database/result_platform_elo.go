package database

type RESULT_PLATFORM_ELO struct {
	ResultID        int64  `json:"result_id" db:"result_id"`
	PlatformEloID   int64  `json:"platform_elo_id" db:"platform_elo_id"`
	VerificationKey string `json:"verification_key" db:"verification_key"`
}

func SelectResultPlatformElo(result_id int64, platform_elo_id int64) (*RESULT_PLATFORM_ELO, error) {
	sth, err := connection.Preparex("SELECT * FROM RESULT_PLATFORM_ELO WHERE result_id = ? AND platform_elo_id = ?")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	var result_platform_elo RESULT_PLATFORM_ELO
	err = sth.QueryRowx(result_id, platform_elo_id).StructScan(&result_platform_elo)

	if err != nil {
		return nil, err
	}

	return &result_platform_elo, nil
}

func SelectResultPlatformElos(where string) (*[]RESULT_PLATFORM_ELO, error) {

	var result_platform_elos []RESULT_PLATFORM_ELO

	rows, err := connection.Queryx("SELECT * FROM RESULT_PLATFORM_ELO " + where)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var result_platform_elo RESULT_PLATFORM_ELO
		err = rows.StructScan(&result_platform_elo)

		if err != nil {
			return nil, err
		}

		result_platform_elos = append(result_platform_elos, result_platform_elo)

	}

	return &result_platform_elos, rows.Err()
}

func InsertResultPlatformElo(result_platform_elo RESULT_PLATFORM_ELO) (*int64, error) {

	sth, err := connection.Prepare("INSERT INTO RESULT_PLATFORM_ELO VALUES (?, ?, ?)")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	result, err := sth.Exec(result_platform_elo.ResultID, result_platform_elo.PlatformEloID, result_platform_elo.VerificationKey)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func ModifyResultPlatformElo(result_id int64, platform_elo_id int64, result_platform_elo RESULT_PLATFORM_ELO) error {

	if len(result_platform_elo.VerificationKey) > 0 {

		sth, err := connection.Prepare("UPDATE RESULT_PLATFORM_ELO SET verification_key = ? WHERE result_id = ? AND platform_elo_id = ?")

		if err != nil {
			return err
		}

		defer sth.Close()

		_, err = sth.Exec(result_id, platform_elo_id, result_platform_elo.VerificationKey)

		if err != nil {
			return err
		}

	}

	return nil
}

func DeleteResultPlatformElo(result_id int64, platform_elo_id int64) error {

	sth, err := connection.Prepare("DELETE FROM RESULT_PLATFORM_ELO WHERE result_id = ? AND platform_elo_id = ?")

	if err != nil {
		return err
	}

	defer sth.Close()

	_, err = sth.Exec(result_id, platform_elo_id)

	if err != nil {
		return err
	}

	return nil
}
