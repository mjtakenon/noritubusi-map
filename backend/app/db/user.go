package db

func (d *DB) InsertUser(userID, hashedPassword string) error {
	query := `INSERT INTO users (id,hashed_password) VALUES (?,?)`

	_, err := d.DB.Exec(query, userID, hashedPassword)

	return err
}
