package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(driverName, connStr string) error {
	var err error
	db, err = sql.Open(driverName, connStr)
	return err
}

func Insert(queryString string) error {
	_, err := db.Exec(queryString)
	return err
}
