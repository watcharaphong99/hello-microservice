package itemHandler

import "github.com/bunkai888/hello-microservice/module/item/itemUsecase"

type (
	itemGrpcHandler struct {
		itemUsecase itemUsecase.ItemUseCaseService
	}
)

func NewItemGrpcHandler(itemUsecase itemUsecase.ItemUseCaseService) *itemGrpcHandler {
	return &itemGrpcHandler{
		itemUsecase: itemUsecase,
	}
}
