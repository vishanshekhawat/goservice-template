package service

import "context"

func (us *UsersService) GetUsers(ctx context.Context) []string {

	return []string{"users1", "users2"}

}
