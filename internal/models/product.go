package models

type Product struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	PriceInCents int    `json:"priceInCents"`
	CreatedAt    string `json:"createdAt"`
}
