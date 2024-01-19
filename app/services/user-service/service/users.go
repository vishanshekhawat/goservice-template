package service

import "context"

func (us *UsersService) GetUsers(ctx context.Context) []string {
	us.log.Infow(ctx, "test")
	return []string{"users1", "users2"}
}
