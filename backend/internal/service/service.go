package service

import (
	"video/backend/dao"

	"google.golang.org/grpc"
)

type Service struct {
	dao *dao.Dao
	RPC *grpc.Server
}

func NewGRPCServer() *grpc.Server {
	s := grpc.NewServer()
	return s
}

func NewService() *Service {
	s := &Service{
		dao: dao.NewDao(),
		RPC: NewGRPCServer(),
	}
	return s
}
