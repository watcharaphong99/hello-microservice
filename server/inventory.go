package server

import (
	"github.com/bunkai888/hello-microservice/module/inventory/inventoryHandler"
	"github.com/bunkai888/hello-microservice/module/inventory/inventoryRepository.go"
	"github.com/bunkai888/hello-microservice/module/inventory/inventoryUsecase"
)

func (s *server) InventoryService() {
	repo := inventoryRepository.NewInventoryRepository(s.db)
	usecase := inventoryUsecase.NewInventoryUsecase(repo)
	httpHandler := inventoryHandler.NewInvetoryHttpHandler(s.cfg, usecase)
	grpcHandler := inventoryHandler.NewInventoryGrpcHandler(usecase)
	queueHandler := inventoryHandler.NewInventoryQueueHanlder(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	inventory := s.app.Group("/inventory_v1")

	inventory.GET("", s.healthCheckService)
}
