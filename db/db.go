package db

import (
	"database/sql"
)

func InitDB(driverName string) error {
	sql.Open(driverName, connectionString)
}
