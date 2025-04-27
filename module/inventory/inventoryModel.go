package inventory

import (
	"github.com/bunkai888/hello-microservice/models"
	"github.com/bunkai888/hello-microservice/module/item"
)

type (
	UpdateInventoryReq struct {
		PlayerId string `json:"player_id" validate:"required,max=64"`
		ItemId   string `json:"item_id" validate:"required,max=64"`
	}

	ItemInventory struct {
		Inventory string `json:"inventory_id"`
		*item.ItemShowCase
	}

	PlayerInventory struct {
		PlayerId string `json:"player_id"`
		*models.PaginateRes
	}
)
