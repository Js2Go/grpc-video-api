package service

import (
	"video/backend/dao"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type Service struct {
	dao    *dao.Dao
	Router *gin.Engine
	RPC    *grpc.Server
}

func NewGRPCServer() *grpc.Server {
	s := grpc.NewServer()
	return s
}

func NewService() *Service {
	s := &Service{
		dao:    dao.NewDao(),
		Router: gin.Default(),
		RPC:    NewGRPCServer(),
	}
	return s
}
