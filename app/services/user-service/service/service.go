package service

import (
	"context"

	"github.com/vishn007/go-service-template/foundation/logger"
)

type Service interface {
	GetUsers(context.Context) []string
}

type UsersService struct {
	log *logger.Logger
}

func NewService(log *logger.Logger) Service {
	return &UsersService{
		log: log,
	}
}
