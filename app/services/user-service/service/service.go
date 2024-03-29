package service

import (
	"context"

	"github.com/vishn007/go-service-template/buisness/repo/userrepo"
	models "github.com/vishn007/go-service-template/buisness/repo/userrepo/model"
	"github.com/vishn007/go-service-template/foundation/logger"
)

type Service interface {
	GetUsers(context.Context) ([]models.User, error)
	CreateUser(context.Context, models.User) (int, error)
}

type UsersService struct {
	repo userrepo.UserRepository
	log  *logger.Logger
}

func NewService(log *logger.Logger, userRepo userrepo.UserRepository) Service {

	//Init Repositories
	return &UsersService{
		log:  log,
		repo: userRepo,
	}
}
