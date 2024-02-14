package grpc

import (
	"context"
	"time"
)

type ctxKey string

const TraceKey ctxKey = "trace_key"

// Values represent state for each request.
type Values struct {
	TraceID       string
	CoRealationID string
	Now           time.Time
	StatusCode    int
}

// GetValues returns the values from the context.
func GetValues(ctx context.Context) *Values {
	v, ok := ctx.Value(TraceKey).(*Values)
	if !ok {
		return &Values{
			TraceID: "00000000-0000-0000-0000-000000000000",
			Now:     time.Now(),
		}
	}

	return v
}

// GetTraceID returns the trace id from the context.
func GetTraceID(ctx context.Context) string {
	v, ok := ctx.Value(TraceKey).(*Values)
	if !ok {
		return "00000000-0000-0000-0000-000000000000"
	}
	return v.TraceID
}

// GetCoRelationID returns the corelation id from the context.
func GetCoRelationID(ctx context.Context) string {
	v, ok := ctx.Value(TraceKey).(*Values)
	if !ok {
		return "00000000-0000-0000-0000-000000000000"
	}
	return v.CoRealationID
}

// GetTime returns the time from the context.
func GetTime(ctx context.Context) time.Time {
	v, ok := ctx.Value(TraceKey).(*Values)
	if !ok {
		return time.Now()
	}
	return v.Now
}

// SetStatusCode sets the status code back into the context.
func SetStatusCode(ctx context.Context, statusCode int) {
	v, ok := ctx.Value(TraceKey).(*Values)
	if !ok {
		return
	}

	v.StatusCode = statusCode
}
