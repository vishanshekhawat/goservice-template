package handlers

import (
	"net/http"
	"os"

	"github.com/vishn007/go-service-template/app/services/user-service/handlers/v1/users"
	"github.com/vishn007/go-service-template/buisness/middleware"
	"github.com/vishn007/go-service-template/foundation/logger"
	"github.com/vishn007/go-service-template/foundation/web"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *logger.Logger
}

func APIMux(cfg APIMuxConfig) *web.App {

	app := web.NewApp(cfg.Shutdown, middleware.Logger(cfg.Log))

	userHandlers := users.New(cfg.Log)
	app.Handle(http.MethodGet, "/test", userHandlers.Test)

	return app
}
