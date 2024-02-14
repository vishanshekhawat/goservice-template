package repo

import (
	"database/sql"
	"fmt"
	"log"

	models "github.com/vishn007/go-service-template/buisness/repo/userrepo/model"
)

type Database interface {
	Connect(models.Config) error
	GetConn() *sql.DB
	Close() error
}

func GetDataBaseConnection(dBCfg models.Config) (Database, error) {
	var database Database

	switch dBCfg.Type {
	case "MYSQL":
		database = &MySQLDB{}
	case "INMEMORY":
		database = &UserCache{}
	default:
		return nil, fmt.Errorf("unsupported database type: %v", dBCfg)
	}

	if err := database.Connect(dBCfg); err != nil {
		return nil, err
	}

	err := database.GetConn().Ping()
	if err != nil {
		log.Fatal("Error:", err)
	}

	return database, nil
}
