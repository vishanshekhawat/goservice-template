package users

import (
	"context"
	"errors"
	"math/rand"
	"net/http"

	"github.com/vishn007/go-service-template/buisness/customerrors"
	"github.com/vishn007/go-service-template/foundation/logger"
	"github.com/vishn007/go-service-template/foundation/web"
)

// Handlers manages the set of user endpoints.
type UserHandlers struct {
	log *logger.Logger
}

// New constructs a handlers for route access.
func New(log *logger.Logger) *UserHandlers {
	return &UserHandlers{
		log: log,
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

	status := struct {
		Status string
	}{
		Status: "OK OK",
	}

	h.log.Infow(ctx, "Test Message")

	return web.Respond(ctx, w, status, http.StatusOK)
}
