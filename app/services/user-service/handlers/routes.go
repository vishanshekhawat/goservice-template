package handlers

import (
	"net/http"
	"os"

	authHandler "github.com/vishn007/go-service-template/app/services/user-service/handlers/v1/auth"
	"github.com/vishn007/go-service-template/app/services/user-service/handlers/v1/users"
	"github.com/vishn007/go-service-template/app/services/user-service/service"
	"github.com/vishn007/go-service-template/buisness/repo"
	"github.com/vishn007/go-service-template/buisness/repo/userrepo"
	"github.com/vishn007/go-service-template/buisness/web/auth"
	"github.com/vishn007/go-service-template/buisness/web/middleware"
	"github.com/vishn007/go-service-template/foundation/logger"
	"github.com/vishn007/go-service-template/foundation/web"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *logger.Logger
	Auth     *auth.Auth
	Db       repo.Database
}

func APIMux(cfg APIMuxConfig) *web.App {

	app := web.NewApp(cfg.Shutdown, middleware.Logger(cfg.Log), middleware.Errors(cfg.Log), middleware.Panics(), middleware.RateLimiter(), middleware.Metrics())

	userRepo := userrepo.GetUserDBRepository(cfg.Db)

	userService := service.NewService(cfg.Log, userRepo)
	userHandlers := users.New(cfg.Log, userService)

	authHandlers := authHandler.New(cfg.Log, cfg.Auth)
	app.Handle(http.MethodGet, "/test", userHandlers.Test)
	app.Handle(http.MethodPost, "/generate-token", authHandlers.GenerateToken)
	app.Handle(http.MethodPost, "/test/auth", userHandlers.Test, middleware.Authenticate(cfg.Auth), middleware.Authorize(cfg.Auth, auth.RuleAdminOnly))
	app.Handle(http.MethodPost, "/api/v1/get-users", userHandlers.GetUsers)

	return app
}
