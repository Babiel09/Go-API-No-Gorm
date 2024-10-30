package models

import "github.com/Babiel09/database"

func Delete(id uint) (int64, error) {
	coon, err := database.DatabaseConnection()
	if err != nil {
		return 0, err
	}

	defer coon.Close()

	res, err := coon.Exec(`DELETE FROM people WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
