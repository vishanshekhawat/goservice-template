package user

import (
	"github.com/vishn007/go-service-template/buisness/repo"
	userdb "github.com/vishn007/go-service-template/buisness/repo/user/mysql"
)

type UserRepository interface {
	CreateUser(name, email string) error
	GetUser(id int) (string, string, error)
	UpdateUser(id int, name, email string) error
	DeleteUser(id int) error
}

func GetUserRepository(dbSql repo.Database) UserRepository {
	return &userdb.UserDB{
		DB: dbSql.GetConn(),
	}
}
