package utils

import (
	"context"
	"errors"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/arashn0uri/go-server/internal/config"
	"github.com/arashn0uri/go-server/internal/constants"
	"github.com/arashn0uri/go-server/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ToPgUUID(id string) (pgtype.UUID, error) {
	u, err := uuid.Parse(id)
	if err != nil {
		return pgtype.UUID{}, err
	}

	return pgtype.UUID{
		Bytes: u,
		Valid: true,
	}, nil
}

func GetUserID(ctx context.Context) (pgtype.UUID, bool) {
	v, ok := ctx.Value(constants.ContextKeyUserID).(pgtype.UUID)
	return v, ok
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
