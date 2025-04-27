package payment

type (
	ItemServiceReq struct {
		Items []*ItemServiceReqDatatum `json:"items" validate:"required"`
	}

	ItemServiceReqDatatum struct {
		ItemId string  `json:"item_id" validate:"require"`
		Price  float64 `json:"price"`
	}
)
