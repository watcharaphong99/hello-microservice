package authHandler

import "github.com/bunkai888/hello-microservice/module/auth/authUsecase"

type (
	authGrpcHandler struct {
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsecaseService) authUsecase.AuthUsecaseService {
	return &authGrpcHandler{authUsecase}
}
