package service

func (s *Service) InitRoutes() {
	grpcRoutes := s.Router.Group("/grpc")
	{
		grpcRoutes.GET("/get_video", s.GRPCGetVideoInfo)
		grpcRoutes.POST("/create_video", s.GRPCCreateVideo)
	}
}
