package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
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

	// make migrations
	if settings.Reload {
		log.Printf("Start reloading database \n")
		err := goose.DownTo(DB, ".", 0)
		if err != nil {
			return err
		}
	}
	log.Printf("Start migrating database \n")
	err = goose.Up(DB, ".")
	if err != nil {
		return err
	}

	return nil
}
