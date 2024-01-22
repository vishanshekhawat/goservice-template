package service

import (
	"context"

	"github.com/vishn007/go-service-template/buisness/repo/user"
	"github.com/vishn007/go-service-template/foundation/logger"
)

type Service interface {
	GetUsers(context.Context) []string
}

type UsersService struct {
	repo user.UserRepository
	log  *logger.Logger
}

func NewService(log *logger.Logger, userRepo user.UserRepository) Service {

	//Init Repositories
	return &UsersService{
		log:  log,
		repo: userRepo,
	}
}
