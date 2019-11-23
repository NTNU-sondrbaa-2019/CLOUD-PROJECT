package database

type PLATFORM_USER struct {
	UserID          int64  `json:"user_id" db:"user_id"`
	PlatformID      int64  `json:"platform_id" db:"platform_id"`
	AccessToken     string `json:"access_token" db:"access_token"`
	VerificationKey string `json:"verification_key" db:"verification_key"`
}

func SelectPlatformUserByUserID(user_id int64) (*[]PLATFORM_USER, error) {

	var platform_users []PLATFORM_USER

	sth, err := connection.Preparex("SELECT * FROM PLATFORM_USER WHERE user_id = ?")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	rows, err := sth.Queryx(user_id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var platform_user PLATFORM_USER
		err = rows.StructScan(&platform_user)

		if err != nil {
			return nil, err
		}

		platform_users = append(platform_users, platform_user)

	}

	return &platform_users, nil
}

func SelectPlatformUserByPlatformID(platform_id int64) (*[]PLATFORM_USER, error) {

	var platform_users []PLATFORM_USER

	sth, err := connection.Preparex("SELECT * FROM PLATFORM_USER WHERE platform_id = ?")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	rows, err := sth.Queryx(platform_id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var platform_user PLATFORM_USER
		err = rows.StructScan(&platform_user)

		if err != nil {
			return nil, err
		}

		platform_users = append(platform_users, platform_user)

	}

	return &platform_users, nil
}

func InsertPlatformUser(platform_user PLATFORM_USER) (*int64, error) {

	sth, err := connection.Prepare("INSERT INTO PLATFORM_USER VALUES (?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	result, err := sth.Exec(platform_user.UserID, platform_user.PlatformID, platform_user.AccessToken, platform_user.VerificationKey)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func ModifyPlatformUser(user_id int64, platform_id int64, platform_user PLATFORM_USER) error {

	if len(platform_user.AccessToken) > 0 {

		sth, err := connection.Prepare("UPDATE PLATFORM_USER SET access_token = ? WHERE user_id = ? AND platform_id = ?")

		if err != nil {
			return err
		}

		defer sth.Close()

		_, err = sth.Exec(platform_user.AccessToken, user_id, platform_id)

		if err != nil {
			return err
		}

	}

	if len(platform_user.VerificationKey) > 0 {

		sth, err := connection.Prepare("UPDATE PLATFORM_USER SET verification_key = ? WHERE user_id = ? AND platform_id = ?")

		if err != nil {
			return err
		}

		defer sth.Close()

		_, err = sth.Exec(platform_user.VerificationKey, user_id, platform_id)

		if err != nil {
			return err
		}

	}

	return nil
}

func DeletePlatformUser(user_id int64, platform_id int64) error {

	sth, err := connection.Prepare("DELETE FROM PLATFORM_USER WHERE user_id = ? AND platform_id = ?")

	if err != nil {
		return err
	}

	defer sth.Close()

	_, err = sth.Exec(user_id, platform_id)

	if err != nil {
		return err
	}

	return nil
}
