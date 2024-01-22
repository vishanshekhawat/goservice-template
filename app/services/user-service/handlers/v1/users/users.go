package users

import (
	"context"
	"errors"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/vishn007/go-service-template/app/services/user-service/service"
	"github.com/vishn007/go-service-template/buisness/customerrors"
	"github.com/vishn007/go-service-template/foundation/logger"
	"github.com/vishn007/go-service-template/foundation/web"
)

// Handlers manages the set of user endpoints.
type UserHandlers struct {
	log *logger.Logger
	srv service.Service
}

// New constructs a handlers for route access.
func New(log *logger.Logger, userService service.Service) *UserHandlers {
	return &UserHandlers{
		log: log,
		srv: userService,
	}
}

// Handlers manages the set of user endpoints.

// Test is our example route.
func (h *UserHandlers) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	// Validate the data
	// Call into the business layer
	if n := rand.Intn(100); n%2 == 0 {
		return customerrors.NewRequestError(errors.New("TRUSTED ERROR"), http.StatusBadRequest)
	}

	res := h.srv.GetUsers(ctx)

	h.log.Infow(ctx, "Test Message")

	return web.Respond(ctx, w, res, http.StatusOK)
}

// Test is our example route.
func (h *UserHandlers) GetUsers(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	// Validate the data
	var app UserRequest
	if err := web.Decode(r, &app); err != nil {
		return err
	}

	// Call into the business layer
	res := h.srv.GetUsers(ctx)

	// Reponse to Client
	resp := UserResponse{
		Users:      res,
		TotalUsers: strconv.Itoa(len(res)),
	}
	return web.Respond(ctx, w, resp, http.StatusOK)
}
