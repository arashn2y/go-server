package users

import (
	"context"

	"github.com/arashn0uri/go-server/internal/constants"
	"github.com/arashn0uri/go-server/internal/models"
	"github.com/arashn0uri/go-server/internal/repository"
	"github.com/arashn0uri/go-server/internal/routes/permissions"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service interface {
	Users(ctx context.Context) ([]models.User, error)
	GetUserByID(ctx context.Context, id pgtype.UUID) (models.User, error)
}

type service struct {
	repository *repository.Queries
	permission permissions.Service
}

func NewService(db *repository.Queries, permission permissions.Service) Service {
	return &service{
		repository: db,
		permission: permission,
	}
}

func (s *service) Users(ctx context.Context) ([]models.User, error) {
	err := s.permission.CheckPermission(ctx, string(constants.ResourceUsers), string(constants.PermissionRead))

	if err != nil {
		return []models.User{}, err
	}

	users, err := s.repository.GetAllUsers(ctx)

	if err != nil {
		return []models.User{}, err
	}
	result := make([]models.User, len(users))
	for i, user := range users {
		result[i] = models.User{
			ID:       user.ID,
			Email:    user.Email,
			Name:     user.Name,
			RoleID:   user.RoleID,
			IsActive: user.IsActive,
		}
	}
	return result, nil
}

func (s *service) GetUserByID(ctx context.Context, id pgtype.UUID) (models.User, error) {
	err := s.permission.CheckPermission(ctx, string(constants.ResourceUsers), string(constants.PermissionRead))

	if err != nil {
		return models.User{}, err
	}

	user, err := s.repository.GetUserByID(ctx, id)

	if err != nil {
		return models.User{}, err
	}

	return models.User{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		RoleID:   user.RoleID,
		IsActive: user.IsActive,
	}, nil
}
