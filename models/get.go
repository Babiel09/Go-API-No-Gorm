package models

import "github.com/Babiel09/database"

func Get(id uint) (people People, err error) {
	coom, err := database.DatabaseConnection()
	if err != nil {
		return
	}
	defer coom.Close()

	row := coom.QueryRow(`SELECT * FROM people WHERE id=$1`, id)

	err = row.Scan(&people.Id, &people.Createdat, &people.Nome, &people.Sexualidade)
	return
}
