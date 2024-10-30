package models

import (
	"github.com/Babiel09/database"
)

func GetAll() (peoples []People, err error) {
	coon, err := database.DatabaseConnection()
	if err != nil {
		return
	}
	//Fecha a conexão assim que termina o método GET
	defer coon.Close()

	rows, err := coon.Query(`SELECT * FROM people`)

	//Inunda por todos os itens retornados
	for rows.Next() {
		var people People
		err = rows.Scan(&people.Id, &people.Createdat, &people.Nome, &people.Sexualidade)
		if err != nil {
			continue
		}
		peoples = append(peoples, people)
	}

	return

}
