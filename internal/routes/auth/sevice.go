package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/alexedwards/argon2id"
	"github.com/arashn0uri/go-server/internal/config"
	"github.com/arashn0uri/go-server/internal/constants"
	"github.com/arashn0uri/go-server/internal/form"
	"github.com/arashn0uri/go-server/internal/models"
	"github.com/arashn0uri/go-server/internal/repository"
)

type Service interface {
	Register(ctx context.Context, data form.Register) (string, error)
	Login(ctx context.Context, data form.Login) (models.LoginResponse, error)
}

type service struct {
	repository *repository.Queries
}

func NewService(db *repository.Queries) Service {
	return &service{
		repository: db,
	}
}

var DefaultParams = &argon2id.Params{
	Memory:      64 * 1024,
	Iterations:  3,
	Parallelism: 2,
	SaltLength:  16,
	KeyLength:   32,
}

func Hash(password string) (string, error) {
	pepper := config.GetEnv(config.EnvPepper)
	return argon2id.CreateHash(password+pepper, DefaultParams)
}

func Verify(password, encodedHash string) (bool, error) {
	pepper := config.GetEnv(config.EnvPepper)

	return argon2id.ComparePasswordAndHash(password+pepper, encodedHash)
}

func GenerateToken(userID pgtype.UUID, email string, roleId int32) (string, error) {
	secret := []byte(config.GetEnv(config.EnvJWTSecret))

	claims := models.Claims{
		UserID: userID,
		Email:  email,
		RoleID: roleId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ValidateToken(tokenStr string) (*models.Claims, error) {
	secret := []byte(config.GetEnv(config.EnvJWTSecret))

	token, err := jwt.ParseWithClaims(tokenStr, &models.Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*models.Claims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

func (s *service) Register(ctx context.Context, data form.Register) (string, error) {

	hash, hashErr := Hash(data.Password)
	if hashErr != nil {
		return "", hashErr
	}

	user, userErr := s.repository.GetUserByEmail(ctx, data.Email)
	if userErr == nil && user.Email != "" {
		return "", errors.New("email already in use")
	}

	role, roleErr := s.repository.GetRoleByName(ctx, data.Role)
	if roleErr != nil || role.Name == string(constants.RoleSuperAdmin) {
		return "", errors.New("invalid role")
	}

	newUserID, err := s.repository.CreateUser(ctx, repository.CreateUserParams{
		Name:     data.Name,
		Email:    data.Email,
		Password: hash,
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

	return "User registered successfully", nil
}

func (s *service) Login(ctx context.Context, data form.Login) (models.LoginResponse, error) {
	user, err := s.repository.GetUserByEmail(ctx, data.Email)

	if err != nil {
		return models.LoginResponse{}, errors.New("invalid email or password")
	}

	match, verifyErr := Verify(data.Password, user.Password)
	if verifyErr != nil || !match {
		return models.LoginResponse{}, errors.New("invalid email or password")
	}

	userPermissions, permErr := s.repository.GetUserPermissions(ctx, user.ID)

	if permErr != nil {
		return models.LoginResponse{}, errors.New("failed to retrieve user permissions")
	}

	// Generate and return JWT token
	jwtToken, tokenErr := GenerateToken(user.ID, user.Email, user.RoleID)
	if tokenErr != nil {
		return models.LoginResponse{}, errors.New("failed to generate JWT token")
	}

	return models.LoginResponse{
		Token:       jwtToken,
		Permissions: userPermissions,
		ID:          user.ID,
	}, nil

}
