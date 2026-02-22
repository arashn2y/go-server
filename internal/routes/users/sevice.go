package users

import (
	"context"

	"github.com/arashn0uri/go-server/internal/repository"
)

type Service interface {
	Users(ctx context.Context) ([]repository.User, error)
	GetUserByEmail(ctx context.Context, email string) (repository.User, error)
}

type service struct {
	repository *repository.Queries
}

func NewService(db *repository.Queries) Service {
	return &service{
		repository: db,
	}
}

func (s *service) Users(ctx context.Context) ([]repository.User, error) {
	users, err := s.repository.GetAllUsers(ctx)

	if err != nil {
		return []repository.User{}, err
	}

	return users, nil
}

func (s *service) GetUserByEmail(ctx context.Context, email string) (repository.User, error) {
	user, err := s.repository.GetUserByEmail(ctx, email)

	if err != nil {
		return repository.User{}, err
	}

	return user, nil
}
