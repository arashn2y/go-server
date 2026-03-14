package permissions

import (
	"context"
	"errors"

	"github.com/arashn0uri/go-server/internal/repository"
	"github.com/arashn0uri/go-server/internal/utils"
)

type Service interface {
	CheckPermission(ctx context.Context, resource, action string) error
}

type service struct {
	repository *repository.Queries
}

func NewService(db *repository.Queries) Service {
	return &service{
		repository: db,
	}
}

func (s *service) CheckPermission(ctx context.Context, resource, action string) error {
	userID, ok := utils.GetUserID(ctx)
	if !ok {
		return errors.New("unauthorized")
	}

	permissions, err := s.repository.GetUserPermissions(ctx, userID)
	if err != nil {
		return errors.New("forbidden")
	}

	for _, p := range permissions {
		if p.Subject == resource && (p.Action == action || p.Action == "all") {
			return nil
		}
	}

	return errors.New("forbidden")
}
