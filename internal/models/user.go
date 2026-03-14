package models

import "github.com/jackc/pgx/v5/pgtype"

type User struct {
	ID       pgtype.UUID `json:"id"`
	Email    string      `json:"email"`
	Name     string      `json:"name"`
	RoleID   int32       `json:"role_id"`
	IsActive bool        `json:"is_active"`
}
