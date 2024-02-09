package repo

import (
	"database/sql"

	models "github.com/vishn007/go-service-template/buisness/repo/userrepo/model"
)

type UserCache struct {
}

// Connect initializes a connection to the in-memory cache.
func (p *UserCache) Connect(cfg models.Config) error {
	return nil
}

// Close closes the connection to the the in-memory cache.
func (p *UserCache) Close() error {

	return nil
}

// Close closes the connection to the in-memory cache.
func (p *UserCache) GetConn() *sql.DB {
	return nil
}
