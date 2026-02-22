package auth

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/arashn0uri/go-server/internal/config"
	"github.com/arashn0uri/go-server/internal/form"
	"github.com/arashn0uri/go-server/internal/models"
	"github.com/arashn0uri/go-server/internal/repository"
	"golang.org/x/crypto/argon2"
)

type Service interface {
	Register(ctx context.Context, data form.Register) (string, error)
	Login(ctx context.Context, data form.Login) (string, error)
}

type service struct {
	repository *repository.Queries
}

func NewService(db *repository.Queries) Service {
	return &service{
		repository: db,
	}
}

var DefaultParams = &models.HashParams{
	Memory:      64 * 1024,
	Iterations:  3,
	Parallelism: 2,
	SaltLength:  16,
	KeyLength:   32,
}

func generateSalt(length uint32) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hash(password string) (string, error) {
	params := DefaultParams

	pepper := config.GetEnv(config.EnvPepper)

	salt, err := generateSalt(params.SaltLength)
	if err != nil {
		return "", err
	}

	// Add pepper before hashing
	combined := password + pepper

	hash := argon2.IDKey(
		[]byte(combined),
		salt,
		params.Iterations,
		params.Memory,
		params.Parallelism,
		params.KeyLength,
	)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encoded := fmt.Sprintf(
		"$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		params.Memory,
		params.Iterations,
		params.Parallelism,
		b64Salt,
		b64Hash,
	)

	return encoded, nil
}

func verify(password, encodedHash string) (bool, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false, errors.New("invalid hash format")
	}

	var memory uint32
	var iterations uint32
	var parallelism uint8

	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d",
		&memory, &iterations, &parallelism)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	pepper := config.GetEnv(config.EnvPepper)

	combined := password + pepper

	computedHash := argon2.IDKey(
		[]byte(combined),
		salt,
		iterations,
		memory,
		parallelism,
		uint32(len(expectedHash)),
	)

	if subtle.ConstantTimeCompare(expectedHash, computedHash) == 1 {
		return true, nil
	}

	return false, nil
}

func (s *service) Register(ctx context.Context, data form.Register) (string, error) {

	hash, hashErr := hash(data.Password)
	if hashErr != nil {
		return "", hashErr
	}

	user, userErr := s.repository.GetUserByEmail(ctx, data.Email)
	if userErr == nil && user.Email != "" {
		return "", errors.New("email already in use")
	}

	err := s.repository.CreateUser(ctx, repository.CreateUserParams{
		Name:     data.Name,
		Email:    data.Email,
		Password: hash,
	})

	if err != nil {
		return "", err
	}

	return "User registered successfully", nil
}

func (s *service) Login(ctx context.Context, data form.Login) (string, error) {
	user, err := s.repository.GetUserByEmail(ctx, data.Email)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	match, verifyErr := verify(data.Password, user.Password)
	if verifyErr != nil {
		return "", errors.New("invalid email or password")
	}

	if !match {
		return "", errors.New("invalid email or password")
	}

	return "Login successful", nil
}
