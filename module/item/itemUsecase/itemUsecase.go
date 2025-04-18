package itemUsecase

import (
	"github.com/bunkai888/hello-microservice/module/item/itemRepository"
)

type (
	ItemUseCaseService interface{}

	itemUsecase struct {
		itemRepository itemRepository.ItemRepositoryService
	}
)

func NewItemUsecase(itemRepository itemRepository.ItemRepositoryService) ItemUseCaseService {
	return &itemUsecase{
		itemRepository: itemRepository}
}
