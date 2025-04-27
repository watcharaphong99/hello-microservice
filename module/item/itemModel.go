package item

import "github.com/bunkai888/hello-microservice/models"

type (
	CreateItemReq struct {
		Title    string  `json:"title" validate:"required,max=64"`
		Price    float64 `json:"price" validate:"require"`
		ImageUrl string  `json:"image_url" validate:"require,max=255"`
		Damage   int     `json:"damage" validate:"require"`
	}

	ItemShowCase struct {
		ItemId   string  `json:"titem_id"`
		Title    string  `json:"title"`
		Price    float64 `json:"price"`
		Damage   int     `json:"damage"`
		ImageUrl string  `json:"image_url"`
	}

	ItemSearch struct {
		Title string `json:"title"`
		models.PaginateReq
	}

	ItemUpdateReq struct {
		Title    string  `json:"title" validate:"required,max=64"`
		Price    float64 `json:"price" validate:"require"`
		ImageUrl string  `json:"image_url" validate:"require,max=255"`
		Damage   int     `json:"damage" validate:"require"`
	}
)
