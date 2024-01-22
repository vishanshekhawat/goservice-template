package repo

import (
	"database/sql"
	"fmt"
)

type Database interface {
	Connect() error
	Close() error
	GetConn() *sql.DB
}

func GetDataBaseConnection(dBType string) (Database, error) {
	var database Database

	switch dBType {
	case "postgres":
		database = &MySQLDB{}
	case "mysql":
		database = &UserCache{}
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dBType)
	}

	if err := database.Connect(); err != nil {
		return nil, err
	}
	return database, nil
}
