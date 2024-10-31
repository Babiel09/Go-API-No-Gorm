package models

import "github.com/Babiel09/database"

func InsertInformation(people People) (id uint, err error) {

	coon, err := database.DatabaseConnection()
	if err != nil {
		return
	}
	defer coon.Close()

	sql := `INSERT INTO pessoas(created_at,nome,sexualidade) VALUES ($1,$2,$3) RETURNING id`
	err = coon.QueryRow(sql, people.Createdat, people.Nome, people.Sexualidade).Scan(&id)
	return
}
