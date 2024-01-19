package auth

import "github.com/vishn007/go-service-template/buisness/validate"

type GenerateTokenRequest struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Validate checks the data in the model is considered clean.
func (app GenerateTokenRequest) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}
	return nil
}

type GenerateTokenResponse struct {
	Token string `json:"token"`
}
