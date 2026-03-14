package auth

import (
	"context"
	"errors"

	"github.com/arashn0uri/go-server/internal/constants"
	"github.com/arashn0uri/go-server/internal/form"
	"github.com/arashn0uri/go-server/internal/models"
	"github.com/arashn0uri/go-server/internal/repository"
	"github.com/arashn0uri/go-server/internal/utils"
)

type Service interface {
	Register(ctx context.Context, data form.Register) (string, error)
	Login(ctx context.Context, data form.Login) (*models.Login, error)
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

	hashedPass, err := utils.Hash(data.Password)
	if err != nil {
		return "", err
	}

	_, err = s.repository.GetUserByEmail(ctx, data.Email)
	if err == nil {
		return "", errors.New("email already in use")
	}

	role, err := s.repository.GetRoleByName(ctx, data.Role)

	if err != nil || role.Name == string(constants.RoleSuperAdmin) {
		return "", errors.New("invalid role")
	}

	newUserID, err := s.repository.CreateUser(ctx, repository.CreateUserParams{
		Name:     data.Name,
		Email:    data.Email,
		Password: hashedPass,
		RoleID:   role.ID,
	})

	if err != nil {
		return "", err
	}

	err = s.repository.AssignRoleToUser(ctx, repository.AssignRoleToUserParams{
		UserID: newUserID,
		RoleID: role.ID,
	})

	if err != nil {
		return "", err
	}

	return "user registered successfully", nil
}

func (s *service) Login(ctx context.Context, data form.Login) (*models.Login, error) {
	user, err := s.repository.GetUserByEmail(ctx, data.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	isVerified, err := utils.Verify(data.Password, user.Password)
	if err != nil || !isVerified {
		return nil, errors.New("invalid email or password")
	}

	userPermissions, err := s.repository.GetUserPermissions(ctx, user.ID)
	if err != nil {
		return nil, errors.New("failed to retrieve user permissions")
	}

	// Generate and return JWT token
	jwtToken, err := utils.GenerateToken(user.ID, user.Email, user.RoleID)
	if err != nil {
		return nil, errors.New("failed to generate JWT token")
	}

	return &models.Login{
		Token:       jwtToken,
		Permissions: userPermissions,
		ID:          user.ID,
	}, nil

}
