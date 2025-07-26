package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB(dataSourceName string) error {
	var err error

	DB, err = sql.Open("sqlite", dataSourceName)
	if err != nil {
		return err
	}

	return DB.Ping()
}
