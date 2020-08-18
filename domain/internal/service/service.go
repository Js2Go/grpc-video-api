package service

import (
	pb "video/proto"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type Service struct {
	Router *gin.Engine
	RPC    pb.VideoClient
}

func NewService() *Service {
	s := &Service{
		Router: gin.Default(),
		RPC:    NewGRPC(),
	}
	return s
}

func NewGRPC() pb.VideoClient {
	Address := "127.0.0.1:23332"
	conn, err := grpc.Dial(Address, grpc.WithInsecure())

	if err != nil {
		grpclog.Fatal(err)
		logrus.Errorf("Start GRPC Error: %s", err.Error())
	}

	c := pb.NewVideoClient(conn)
	return c
}
