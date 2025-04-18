package itemHandler

import (
	"github.com/bunkai888/hello-microservice/config"
	"github.com/bunkai888/hello-microservice/module/item/itemUsecase"
)

type (
	ItemHttpHandlerService interface{}

	itemHttpHandler struct {
		cfg         *config.Config
		itemUsecase itemUsecase.ItemUseCaseService
	}
)

func NewItemHttpHandler(cfg *config.Config, itemUsecase itemUsecase.ItemUseCaseService) ItemHttpHandlerService {
	return &itemHttpHandler{
		cfg:         cfg,
		itemUsecase: itemUsecase,
	}
}
