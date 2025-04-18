package middlewareHandler

import (
	"github.com/bunkai888/hello-microservice/config"
	"github.com/bunkai888/hello-microservice/module/middleware/middlewareUsecase"
)

type (
	MiddlewareHandlerService interface{}

	middlewareHandler struct {
		cfg               *config.Config
		middlewareUsecase middlewareUsecase.MiddlewareUsecaseService
	}
)

func NewMiddlewareHandler(cfg *config.Config, middlewareUsecase middlewareUsecase.MiddlewareUsecaseService) MiddlewareHandlerService {
	return &middlewareHandler{
		cfg:               cfg,
		middlewareUsecase: middlewareUsecase,
	}
}
