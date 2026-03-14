package models

import "github.com/jackc/pgx/v5/pgtype"

type Product struct {
	ID          pgtype.UUID `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	ImageUrl    string      `json:"imageUrl"`
	Price       float64     `json:"price"`
}
