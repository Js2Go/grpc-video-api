package server

import (
	"net"
	"video/backend/internal/service"

	pb "video/proto"

	"github.com/sirupsen/logrus"
)

func NewGRPCServer(s *service.Service) {
	Address := "127.0.0.1:23332"
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		logrus.Errorf("Listen Error: %s", err.Error())
		panic(err)
	}

	pb.RegisterVideoServer(s.RPC, s)
	s.RPC.Serve(listen)
}
