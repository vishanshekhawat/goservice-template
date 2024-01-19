package middleware

import (
	"context"
	"errors"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"

	"github.com/vishn007/go-service-template/buisness/customerrors"
	"github.com/vishn007/go-service-template/foundation/web"
)

func RateLimiter() web.Middleware {

	type client struct {
		limiter  *rate.Limiter
		lastSeen time.Time
	}
	var (
		mu      sync.Mutex
		clients = make(map[string]*client)
	)
	go func() {
		for {
			time.Sleep(time.Minute)
			// Lock the mutex to protect this section from race conditions.
			mu.Lock()
			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()

	m := func(handler web.Handler) web.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				return err
			}
			// Lock the mutex to protect this section from race conditions.
			mu.Lock()
			if _, found := clients[ip]; !found {
				clients[ip] = &client{limiter: rate.NewLimiter(1, 1)}
			}
			clients[ip].lastSeen = time.Now()
			if !clients[ip].limiter.Allow() {
				mu.Unlock()

				err := customerrors.RateLimitError{
					Err:    errors.New("The API is at capacity, try again later."),
					Status: http.StatusTooManyRequests,
				}

				return &err
			}
			mu.Unlock()
			err = handler(ctx, w, r)
			return err
		}
		return h
	}

	return m
}
