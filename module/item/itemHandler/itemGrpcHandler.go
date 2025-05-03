package itemHandler

import (
	"context"

	itemPb "github.com/bunkai888/hello-microservice/module/item/itemPb"
	"github.com/bunkai888/hello-microservice/module/item/itemUsecase"
)

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

func (g *itemGrpcHandler) FindItemsInIds(ctx context.Context, req *itemPb.FindItemsInIdsReq) (*itemPb.FindItemsInIdsRes, error) {
	return nil, nil
}
