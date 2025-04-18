package inventoryHandler

import (
	"github.com/bunkai888/hello-microservice/config"
	"github.com/bunkai888/hello-microservice/module/inventory/inventoryUsecase"
)

type (
	InventoryHttpHandler interface{}

	inventoryHttpHandler struct {
		cfg              *config.Config
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInvetoryHttpHandler(cfg *config.Config, inventoryUsecase inventoryUsecase.InventoryUsecaseService) InventoryHttpHandler {
	return &inventoryHttpHandler{
		cfg:              cfg,
		inventoryUsecase: inventoryUsecase,
	}
}
