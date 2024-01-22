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
}

func GetUserDBRepository(dbSql repo.Database) UserRepository {
	return &userdb.UserDB{
		DB: dbSql.GetConn(),
	}
}

func GetUserCacheRepository() UserRepository {
	return &cachedb.CacheDB{
		Users: map[uuid.UUID]models.User{},
	}
}
