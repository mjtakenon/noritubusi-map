package db

import "time"

type UserInfo struct {
	Id             string    `db:"id"`
	HashedPassword string    `db:"hashed_password" json:"-"`
	CreateTime     time.Time `db:"create_time"`
	ChangeTime     time.Time `db:"change_time"`
}

func (d *DB) InsertUser(userID, hashedPassword string) error {
	query := `INSERT INTO users (id,hashed_password) VALUES (?,?)`

	_, err := d.DB.Exec(query, userID, hashedPassword)

	return err
}

func (d *DB) GetUserInfoByUserID(userID string) (UserInfo, error) {
	query := `SELECT * FROM users WHERE id = ?`

	var info UserInfo
	err := d.DB.Get(&info, query, userID)

	return info, err
}

func (d *DB) UpdateUser(userID, newHashedPassword string) error {
	query := `UPDATE users SET hashed_password = ? WHERE id = ?`

	_, err := d.DB.Exec(query, newHashedPassword, userID)

	return err
}

func (d *DB) DeleteUser(userID string) error {
	query := `DELETE FROM users WHERE id = ?`

	_, err := d.DB.Exec(query, userID)

	return err
}
