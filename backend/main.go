package main

import (
	"os"
	"os/signal"
	"syscall"
	"video/backend/internal/server"
	"video/backend/internal/service"
)

func main() {
	s := service.NewService()
	server.NewHttpServer(s)
	server.NewGRPCServer(s)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		switch <-c {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			return
		case syscall.SIGHUP:
		}
	}
}
