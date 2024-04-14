package server

import (
	"banners-service/internal/config"
	"net/http"
)

type Server struct {
	standardServer *http.Server
}

func New(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		standardServer: &http.Server{
			Addr:    ":" + cfg.HTTP.Port,
			Handler: handler,
		},
	}
}

func (s *Server) Run() error {
	return s.standardServer.ListenAndServe()
}
