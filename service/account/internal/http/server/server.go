package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo *echo.Echo
	Port string
}

func New(port string) Server {
	e := echo.New()

	return Server{Echo: e, Port: fmt.Sprintf(":%s", port)}
}

func (s *Server) Start() error {
	// Creating new context for waiting shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Starting server
	go func() {
		s.Echo.Logger.Infof("Starting on %d", s.Port)
		if err := s.Echo.Start(s.Port); err != nil && err != http.ErrServerClosed {
			s.Echo.Logger.Errorf("HTTP server ListenAndServe: %v", err)
		}
	}()

	// Waiting done signal
	<-ctx.Done()

	return s.GracefulStop()
}

func (s *Server) GracefulStop() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Shutdown Echo
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Echo.Shutdown(ctx); err != nil {
		s.Echo.Logger.Fatal(err)
		return err
	}
	return nil
}
