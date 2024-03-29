package web

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/google/uuid"
)

// A Handler is a type that handles a http request within our own little mini
// framework.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct.
type App struct {
	*httptreemux.ContextMux
	shutdown chan os.Signal
	mw       []Middleware
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(shutdown chan os.Signal, mw ...Middleware) *App {
	return &App{
		ContextMux: httptreemux.NewContextMux(),
		shutdown:   shutdown,
		mw:         mw,
	}
}

func (app *App) NewGroup(path string) *httptreemux.ContextGroup {
	return app.NewContextGroup(path)
}

// Handle sets a handler function for a given HTTP method and path pair
// to the application server mux.
func (a *App) Handle(method string, path string, handler Handler, mw ...Middleware) {

	handler = wrapMiddleware(mw, handler)
	handler = wrapMiddleware(a.mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {

		// ANY CODE I WANT

		requestID := r.Header.Get("Request-ID")
		if requestID == "" {
			requestID = uuid.NewString()
		}

		v := Values{
			TraceID:       uuid.NewString(),
			CoRealationID: requestID,
			Now:           time.Now().UTC(),
		}

		ctx := context.WithValue(r.Context(), traceKey, &v)
		if err := handler(ctx, w, r); err != nil {
			return
		}

		// ANY CODE I WANT
	}

	a.ContextMux.Handle(method, path, h)
}
