package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"kodix/src/internal/http/api"
	"kodix/src/internal/logger"
)

type Application struct {
	r *api.Router
}
var (
	network = "tcp4"
	port = os.Getenv("API_PORT")
)

func New() *Application {
	return &Application{
		r: api.NewRouter(),
	}
}

func (a *Application) Run() error {
	addr := fmt.Sprintf(":%s", port)

	a.r.RegisterRoutes()
	s := &http.Server{
		Addr:        addr,
		Handler:     a.r.Handler,
		ReadTimeout: 10 * time.Second,
	}

	listener, err := net.Listen(network, addr)
	if err != nil {
		logger.Logger.Fatalf("Failed to listen and serve: %+v", err)
	}
	go func() {
		logger.Logger.Infof("Listening http server on %s...\n", addr)

		if err := s.Serve(listener); err != nil {
			logger.Logger.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return s.Shutdown(ctx)
}
