package server

import (
	"net/http"
	"video/domain/internal/service"

	"github.com/sirupsen/logrus"
)

func NewServer(s *service.Service) {
	s.InitRoutes()
	server := http.Server{
		Addr:    ":22222",
		Handler: s.Router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			logrus.Errorf("HTTP Server Error: %s", err.Error())
		}
	}()
}
