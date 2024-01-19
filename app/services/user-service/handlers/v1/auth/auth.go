package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/vishn007/go-service-template/buisness/auth"
	"github.com/vishn007/go-service-template/foundation/logger"
	"github.com/vishn007/go-service-template/foundation/web"
)

// Handlers manages the set of user endpoints.
type AuthHandlers struct {
	log  *logger.Logger
	auth *auth.Auth
}

// New constructs a handlers for route access.
func New(log *logger.Logger, auth *auth.Auth) *AuthHandlers {

	return &AuthHandlers{
		log:  log,
		auth: auth,
	}
}

func (h *AuthHandlers) GenerateToken(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	// Validate the data
	var app GenerateTokenRequest
	if err := web.Decode(r, &app); err != nil {
		return err
	}
	expirationTime := time.Now().Add(time.Minute * time.Duration(15))

	tokenclaims := &auth.Claims{
		UserID: app.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: expirationTime,
			},
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
			Issuer: h.auth.Issuer,
		},
	}
	res, err := h.auth.GenerateToken("KID", tokenclaims)

	if err != nil {
		return err
	}

	resp := GenerateTokenResponse{
		Token: res,
	}

	return web.Respond(ctx, w, resp, http.StatusOK)
}
