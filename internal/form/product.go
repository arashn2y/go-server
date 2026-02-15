package form

type CreateProductRequest struct {
	Name         string `json:"name" validate:"required"`
	Description  string `json:"description"`
	PriceInCents int64  `json:"priceInCents" validate:"required,gt=0"`
}

type UpdateProductRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	PriceInCents int64  `json:"priceInCents"`
}
