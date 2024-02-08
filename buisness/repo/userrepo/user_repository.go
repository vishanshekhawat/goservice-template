package userrepo

import (
	"github.com/google/uuid"
	"github.com/vishn007/go-service-template/buisness/repo"
	"github.com/vishn007/go-service-template/buisness/repo/userrepo/cachedb"
	models "github.com/vishn007/go-service-template/buisness/repo/userrepo/model"
	userdb "github.com/vishn007/go-service-template/buisness/repo/userrepo/mysql"
)

type UserRepository interface {
	CreateUser(name, email string) error
	GetUser(id uuid.UUID) (models.User, error)
	UpdateUser(id uuid.UUID, name, email string) error
	DeleteUser(id uuid.UUID) error
	GetUsers() ([]models.User, error)
}

func GetUserRepository(dbSql repo.Database) UserRepository {

	var dbType string

	switch dbSql.(type) {
	case *repo.MySQLDB:
		dbType = "MYSQL"
	case *repo.UserCache:
		dbType = "IN-MEMORY"
	}

	switch dbType {
	case "MYSQL":
		return &userdb.UserDB{
			DB: dbSql.GetConn(),
		}
	case "IN-MEMORY":
		return &cachedb.CacheDB{
			Users: map[uuid.UUID]models.User{},
		}
	}
	return nil
}
