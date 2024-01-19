package handlers

import (
	"net/http"
	"os"

	authHandler "github.com/vishn007/go-service-template/app/services/user-service/handlers/v1/auth"
	"github.com/vishn007/go-service-template/app/services/user-service/handlers/v1/users"
	"github.com/vishn007/go-service-template/buisness/auth"
	"github.com/vishn007/go-service-template/buisness/middleware"
	"github.com/vishn007/go-service-template/foundation/logger"
	"github.com/vishn007/go-service-template/foundation/web"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *logger.Logger
	Auth     *auth.Auth
}

func APIMux(cfg APIMuxConfig) *web.App {

	app := web.NewApp(cfg.Shutdown, middleware.Logger(cfg.Log), middleware.Errors(cfg.Log), middleware.Panics(), middleware.RateLimiter(), middleware.Metrics())

	userHandlers := users.New(cfg.Log)
	authHandlers := authHandler.New(cfg.Log, cfg.Auth)
	app.Handle(http.MethodGet, "/test", userHandlers.Test)
	app.Handle(http.MethodPost, "/generate-token", authHandlers.GenerateToken)
	app.Handle(http.MethodPost, "/test/auth", userHandlers.Test, middleware.Authenticate(cfg.Auth))
	app.Handle(http.MethodPost, "/api/v1/get-users", userHandlers.GetUsers)

	return app
}
