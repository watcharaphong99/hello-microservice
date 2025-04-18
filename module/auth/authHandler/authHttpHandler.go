package authHandler

import (
	"github.com/bunkai888/hello-microservice/config"
	"github.com/bunkai888/hello-microservice/module/auth/authUsecase"
)

type (
	AuthHttpHandlerService interface{}

	authHttpHandler struct {
		cfg         *config.Config
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthHtppHandler(cfg *config.Config, authUsecase authUsecase.AuthUsecaseService) AuthHttpHandlerService {
	return &authHttpHandler{cfg, authUsecase}
}
