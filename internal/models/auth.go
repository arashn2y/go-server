package models

import (
	"github.com/arashn0uri/go-server/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Claims struct {
	UserID pgtype.UUID `json:"user_id"`
	Email  string      `json:"email"`
	RoleID int32       `json:"role_id"`
	jwt.RegisteredClaims
}

type LoginResponse struct {
	Token       string                             `json:"token"`
	Permissions []repository.GetUserPermissionsRow `json:"permissions"`
}
