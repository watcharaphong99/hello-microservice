package server

import (
	"github.com/bunkai888/hello-microservice/module/auth/authHandler"
	"github.com/bunkai888/hello-microservice/module/auth/authRepository"
	"github.com/bunkai888/hello-microservice/module/auth/authUsecase"
)

func (s *server) authService() {
	repo := authRepository.NewAuthRepository(s.db)
	usecase := authUsecase.NewAuthUsecase(repo)
	httpHandler := authHandler.NewAuthHtppHandler(s.cfg, usecase)
	grpcHandler := authHandler.NewAuthGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth_v1")

	_ = auth
}
