package users

import (
	"context"
	"encoding/json"
	"net/http"
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

	jsonData, err := json.Marshal(status)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}
