package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo       *echo.Echo
	HttpServer *http.Server
}

func New(port string) (Server, error) {
	e := echo.New()

	s := http.Server{
		Addr:    port,
		Handler: e,
		//ReadTimeout: 30 * time.Second, // customize http.Server timeouts
	}

	return Server{Echo: e, HttpServer: &s}, nil
}

func (s *Server) Start() error {

}

func (s *Server) Stop() error
