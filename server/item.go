package server

import (
	"github.com/bunkai888/hello-microservice/module/item/itemHandler"
	"github.com/bunkai888/hello-microservice/module/item/itemRepository"
	"github.com/bunkai888/hello-microservice/module/item/itemUsecase"
)

func (s *server) itemService() {
	repo := itemRepository.NewItemRepository(s.db)
	usecase := itemUsecase.NewItemUsecase(repo)
	httpHander := itemHandler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := itemHandler.NewItemGrpcHandler(usecase)

	_ = httpHander
	_ = grpcHandler

	item := s.app.Group("/item_v1")

	//Health Check
	_ = item
}
