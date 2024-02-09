package service

import (
	"context"

	models "github.com/vishn007/go-service-template/buisness/repo/userrepo/model"
)

func (us *UsersService) GetUsers(ctx context.Context) ([]models.User, error) {

	users, err := us.repo.GetUsers()
	if err != nil {
		return users, err
	}
	return users, nil

}

func (us *UsersService) CreateUser(ctx context.Context, user models.User) (int, error) {

	res, err := us.repo.CreateUser(ctx, user)
	if err != nil {
		return 0, err
	}
	return res, nil

}
