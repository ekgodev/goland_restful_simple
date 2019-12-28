package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var (
	DB *sql.DB
)

func Connect(settings Settings) (err error) {

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s, dbname=%s sslmode=disable",
		settings.Host, settings.Port, settings.User, settings.Pass, settings.Name)

	DB, err = sql.Open("postgres", sqlInfo)
	if err != nil {
		return err
	}

	log.Printf("Database connection was created: %s \n", sqlInfo)
	return nil
}
