package users

import (
	models "github.com/vishn007/go-service-template/buisness/repo/userrepo/model"
	"github.com/vishn007/go-service-template/buisness/validate"
)

type UserResponse struct {
	Users      []models.User `json:"users"`
	TotalUsers string        `json:"total_users"`
}

// AppNewUser contains information needed to create a new user.
type UserRequest struct {
	Token string `json:"token" validate:"required"`
}

// Validate checks the data in the model is considered clean.
func (app UserRequest) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}
	return nil
}
