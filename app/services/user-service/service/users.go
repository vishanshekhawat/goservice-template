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
