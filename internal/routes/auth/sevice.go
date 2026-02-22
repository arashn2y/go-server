package auth

import (
	"context"

	"github.com/arashn0uri/go-server/internal/form"
	"github.com/arashn0uri/go-server/internal/repository"
)

type Service interface {
	Register(ctx context.Context, data form.Register) (string, error)
}

type service struct {
	repository *repository.Queries
}

func NewService(db *repository.Queries) Service {
	return &service{
		repository: db,
	}
}

func (s *service) Register(ctx context.Context, data form.Register) (string, error) {
	err := s.repository.CreateUser(ctx, repository.CreateUserParams{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	})

	if err != nil {
		return "", err
	}

	return "User registered successfully", nil
}
