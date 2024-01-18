package users

import (
	"context"
	"net/http"

	"github.com/vishn007/go-service-template/foundation/web"
)

// Handlers manages the set of user endpoints.

// Test is our example route.
func Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	// Validate the data
	// Call into the business layer

	status := struct {
		Status string
	}{
		Status: "OK OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
