package models

import "github.com/Babiel09/database"

func Update(id uint, peoples People) (int64, error) {
	coon, err := database.DatabaseConnection()
	if err != nil {
		return 0, err
	}
	defer coon.Close()

	res, err := coon.Exec(`UPDATE people SET created_at=$2, nome=$3, sexualidade=$4 WHERE id=$1`, id, peoples.Createdat, peoples.Nome, peoples.Sexualidade)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()

}
