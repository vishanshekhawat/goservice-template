package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/vishn007/go-service-template/foundation/logger"
)

func main() {

	log, err := logger.New("SWAGGER-SERVICE")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	if err := run(log); err != nil {
		log.Errorw(context.TODO(), "startup", "ERROR", err)
		log.Sync()
		os.Exit(1)
	}

}

func run(log *logger.Logger) error {
	//----------------------Service Start-------------------------//
	log.Info(context.TODO(), "starting Swagger UI Service")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	fs := http.FileServer(http.Dir("./doc"))

	http.Handle("/doc/", http.StripPrefix("/doc/", fs))
	http.Handle("/api/", http.StripPrefix("/api/", http.FileServer(http.Dir("./api"))))

	serverErrors := make(chan error, 1)

	go func() {
		log.Infow(context.TODO(), "startup", "status", "api router started", "host", "8090")
		serverErrors <- http.ListenAndServe(":8090", nil)
	}()

	// -------------------------------------------------------------------------
	// Shutdown

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.Infow(context.TODO(), "shutdown", "status", "shutdown started", "signal", sig)
	}

	return nil

}
