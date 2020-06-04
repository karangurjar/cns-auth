package db

import (
	"database/sql"

	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

var db *sql.DB

func GetCnsDB() *sql.DB {
	return db
}

func InitDB(driverName, connStr string) error {
	var err error
	db, err = sql.Open(driverName, connStr)
	if err != nil {
		return err
	}

	//create database if not exists
	checkDBQuery := `SELECT EXISTS(select datname from pg_catalog.pg_database where datname='cns');`
	var exists bool
	err = db.QueryRow(checkDBQuery).Scan(&exists)
	if err != nil {
		log.Errorf("error while creating database, Error: %q", err.Error())
		return err
	}

	if !exists {
		createDBQuery := `CREATE DATABASE cns;`
		_, err = db.Exec(createDBQuery)
		if err != nil {
			return err
		}
	}

	//creat users tables if dosen't exists
	createTableQuery := `CREATE TABLE IF NOT EXISTS users(username varchar(30), password varchar(300), email varchar(50))`
	_, err = db.Exec(createTableQuery)
	return err
}

func Insert(queryString string) error {
	//TODO improve this method with Query
	_, err := db.Exec(queryString)

	return err
}
