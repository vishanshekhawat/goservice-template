package logger

import (
	"context"

	"github.com/vishn007/go-service-template/foundation/web"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

func (log *Logger) Errorw(ctx context.Context, message string, args ...interface{}) {
	args = setCommonArgs(ctx, args)
	log.SugaredLogger.Errorw(message, args...)
}

func (log *Logger) Error(ctx context.Context, args ...interface{}) {
	args = setCommonArgs(ctx, args)
	log.SugaredLogger.Error(args...)
}

func (log *Logger) Infow(ctx context.Context, message string, args ...interface{}) {
	args = setCommonArgs(ctx, args)
	log.SugaredLogger.Infow(message, args...)
}

func (log *Logger) Info(ctx context.Context, args ...interface{}) {
	args = setCommonArgs(ctx, args)
	log.SugaredLogger.Info(args...)
}

func (log *Logger) Warnw(ctx context.Context, message string, args ...interface{}) {
	args = setCommonArgs(ctx, args)
	log.SugaredLogger.Warnw(message, args...)
}

func (log *Logger) Warn(ctx context.Context, args ...interface{}) {
	args = setCommonArgs(ctx, args)
	log.SugaredLogger.Warn(args...)
}

func (log *Logger) Debugw(ctx context.Context, message string, args ...interface{}) {
	args = setCommonArgs(ctx, args)
	log.SugaredLogger.Debugw(message, args...)
}

func (log *Logger) Debug(ctx context.Context, args ...interface{}) {
	args = setCommonArgs(ctx, args)
	log.SugaredLogger.Debug(args...)
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

// setCoRelationID adds request-id key-value to the argument list
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
