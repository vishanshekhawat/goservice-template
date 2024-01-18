package logger

import (
	"context"
	"crypto/md5"
	"fmt"
	"os"
	"strings"

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

func (log *Logger) LogEnvVars(ctx context.Context, secretEnvVars []string) {
	argsList := make([]interface{}, 0)

	for _, envVar := range os.Environ() {
		envVarPair := strings.Split(envVar, "=")
		if len(envVarPair) != 2 {
			continue
		}
		envVarkey := envVarPair[0]
		envVarValue := envVarPair[1]

		for _, s := range secretEnvVars {
			if s == envVarkey {
				envVarValue = fmt.Sprintf("%x", md5.Sum([]byte(envVarValue)))
			}
		}

		argsList = SetKeyValueToArgs(argsList, envVarkey, envVarValue)
	}

	log.Infow("Environment Variables", argsList...)
}

func SetKeyValueToArgs(args []interface{}, key interface{}, val interface{}) []interface{} {
	args = append(args, key, val)
	return args
}
