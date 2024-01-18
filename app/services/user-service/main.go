package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/vishn007/go-service-template/foundation/logger"
	"go.uber.org/zap"
)

var build = "develop"

func main() {
	log, err := logger.New("USER-SERVICE")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	if err := run(log); err != nil {
		log.Errorw("startup", "ERROR", err)
		log.Sync()
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {
	log.Infow("startup", "GOMAXPROCS", runtime.GOMAXPROCS(0), "BUILD - ", build)
	return nil
}
