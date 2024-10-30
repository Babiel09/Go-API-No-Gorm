package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Babiel09/configs"
	_ "github.com/lib/pq"
)

func DatabaseConnection() (*sql.DB, error) {
	dataConf := configs.GetDB()
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dataConf.Host, dataConf.Port, dataConf.User, dataConf.Pass, dataConf.DataBase)

	coon, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	err = coon.Ping()
	return coon, err
}
