package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
	port string
}

func New(port string) *Server {
	s := &Server{
		echo: echo.New(),
		port: fmt.Sprintf(":%s", port),
	}
	s.setupMiddleware()
	s.setupRoutes()
	return s
}

func (s *Server) Start() error {
	// Creating new context for waiting shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Starting server
	go func() {
		s.echo.Logger.Infof("Starting on %d", s.port)
		if err := s.echo.Start(s.port); err != nil && err != http.ErrServerClosed {
			s.echo.Logger.Errorf("HTTP server ListenAndServe: %v", err)
		}
	}()

	// Waiting done signal
	<-ctx.Done()

	return s.gracefulStop()
}

func (s *Server) gracefulStop() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Shutdown Echo
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.echo.Shutdown(ctx); err != nil {
		s.echo.Logger.Fatal(err)
		return err
	}
	return nil
}

func (s *Server) setupMiddleware() {
	s.echo.Use(middleware.Logger())

	s.echo.Use(middleware.Recover())

	// CORS
	s.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))
}

func (s *Server) setupRoutes() {
	// Группа API
	std := s.echo.Group("/api")
	// auth := s.echo.Group()
	// admin := s.echo.Group()

	std.GET("/ping", func(c echo.Context) error { return c.JSON(200, "PONG") })
}
