package logger

import (
	"context"

	"github.com/vishn007/go-service-template/foundation/web"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	RequestIDKey     = "request_id"
	CustomerIDKey    = "customer_id"
	LogRequestIDKey  = "request_id"
	LogCustomerIDKey = "customer_id"
	LogLevelEnvVar   = "LOG_LEVEL"
)

type Logger struct {
	*zap.SugaredLogger
}

// New constructs a Sugared Logger that writes to stdout and
// provides human-readable timestamps.
func New(service string, outputPaths ...string) (*Logger, error) {
	config := zap.NewProductionConfig()

	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	config.InitialFields = map[string]any{
		"service": service,
	}

	config.OutputPaths = []string{"stdout"}
	if outputPaths != nil {
		config.OutputPaths = outputPaths
	}

	log, err := config.Build(zap.WithCaller(true))
	if err != nil {
		return nil, err
	}

	return &Logger{log.Sugar()}, nil
}

func (log *Logger) Infow(ctx context.Context, message string, args ...interface{}) {
	args = setCommonArgs(ctx, args)
	log.SugaredLogger.Infow(message, args...)
}

// setCommonArgs adds common set of key-value pairs to the arguments list
func setCommonArgs(ctx context.Context, args []interface{}) []interface{} {
	args = setTraceID(ctx, args)
	args = setCoRelationID(ctx, args)
	return args
}

// setTraceID adds request-id key-value to the argument list
func setTraceID(ctx context.Context, args []interface{}) []interface{} {
	reqID := web.GetTraceID(ctx)
	if reqID == "00000000-0000-0000-0000-000000000000" {
		return args
	}
	args = SetKeyValueToArgs(args, "trace_id", reqID)
	return args
}

// setTraceID adds request-id key-value to the argument list
func setCoRelationID(ctx context.Context, args []interface{}) []interface{} {
	reqID := web.GetCoRelationID(ctx)
	if reqID == "00000000-0000-0000-0000-000000000000" {
		return args
	}
	args = SetKeyValueToArgs(args, "corelation_id", reqID)
	return args
}

func SetKeyValueToArgs(args []interface{}, key interface{}, val interface{}) []interface{} {
	args = append(args, key, val)
	return args
}
