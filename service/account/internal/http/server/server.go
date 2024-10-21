package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo       *echo.Echo
	HttpServer *http.Server
}

func New(port string) Server {
	e := echo.New()

	s := http.Server{
		Addr:    port,
		Handler: e,
		//ReadTimeout: 30 * time.Second, // customize http.Server timeouts
	}

	return Server{Echo: e, HttpServer: &s}
}

func (s *Server) Start() error {
	// Creating new context for waiting shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Starting server
	go func() {
		if err := s.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Echo.Logger.Errorf("HTTP server ListenAndServe: %v", err)
		}
	}()

	// Waiting done signal
	<-ctx.Done()

	return s.GracefulStop()
}

func (s *Server) GracefulStop() error {
	// Creating new context with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown Echo
	if err := s.Echo.Shutdown(shutdownCtx); err != nil {
		s.Echo.Logger.Errorf("Echo shutdown: %v", err)
		return err
	}

	// Shutdown HttpServer
	if err := s.HttpServer.Shutdown(shutdownCtx); err != nil {
		s.Echo.Logger.Errorf("Server forced shutdown: %v", err)
		return err
	}

	return nil
}
