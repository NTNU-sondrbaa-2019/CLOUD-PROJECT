package database

type GROUP_USER struct {
	GroupID int64 `json:"group_id" db:"group_id"`
	UserID  int64 `json:"user_id" db:"user_id"`
}

func SelectGroupUserByUserID(user_id int64) (*[]GROUP_USER, error) {

	var group_users []GROUP_USER

	sth, err := connection.Preparex("SELECT * FROM GROUP_USER WHERE user_id = ?")

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

		var group_user GROUP_USER
		err = rows.StructScan(&group_user)

		if err != nil {
			return nil, err
		}

		group_users = append(group_users, group_user)

	}

	return &group_users, nil
}

func SelectGroupUserByGroupID(group_id int64) (*[]GROUP_USER, error) {

	var group_users []GROUP_USER

	sth, err := connection.Preparex("SELECT * FROM GROUP_USER WHERE group_id = ?")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	rows, err := sth.Queryx(group_id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var group_user GROUP_USER
		err = rows.StructScan(&group_user)

		if err != nil {
			return nil, err
		}

		group_users = append(group_users, group_user)

	}

	return &group_users, nil
}

func InsertGroupUser(group_user GROUP_USER) (*int64, error) {

	sth, err := connection.Prepare("INSERT INTO GROUP_USER VALUES (?, ?)")

	if err != nil {
		return nil, err
	}

	defer sth.Close()

	result, err := sth.Exec(group_user.GroupID, group_user.UserID)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func DeleteGroupUser(group_id int64, user_id int64) error {

	sth, err := connection.Prepare("DELETE FROM GROUP_USER WHERE group_id = ? AND user_id = ?")

	if err != nil {
		return err
	}

	defer sth.Close()

	_, err = sth.Exec(group_id, user_id)

	if err != nil {
		return err
	}

	return nil
}
