package form

import "mime/multipart"

type CreateProductRequest struct {
	Name         string                `form:"name" validate:"required"`
	Description  string                `form:"description"`
	PriceInCents int64                 `form:"priceInCents" validate:"required,gt=0"`
	Image        multipart.File        `form:"image,file"   validate:"required"`
	ImageHeader  *multipart.FileHeader `form:"image,header"`
}

type UpdateProductRequest struct {
	Name         string                `form:"name"`
	Description  string                `form:"description"`
	PriceInCents int64                 `form:"priceInCents" validate:"omitempty,gt=0"`
	Image        multipart.File        `form:"image,file"`
	ImageHeader  *multipart.FileHeader `form:"image,header"`
}
