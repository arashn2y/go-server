package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvKey string

const (
	EnvAddr               EnvKey = "ADDR"
	EnvDSN                EnvKey = "DSN"
	EnvPepper             EnvKey = "PASSWORD_PEPPER"
	EnvSuperAdminName     EnvKey = "SUPERADMIN_NAME"
	EnvSuperAdminEmail    EnvKey = "SUPERADMIN_EMAIL"
	EnvSuperAdminPassword EnvKey = "SUPERADMIN_PASSWORD"
	EnvJWTSecret          EnvKey = "JWT_SECRET"
	R2AccountID           EnvKey = "R2_ACCOUNT_ID"
	R2AccessKeyID         EnvKey = "R2_ACCESS_KEY_ID"
	R2SecretAccessKey     EnvKey = "R2_SECRET_ACCESS_KEY"
	R2Bucket              EnvKey = "R2_BUCKET"
)

type Config struct {
	Addr            string
	DB              DBConfig
	PASSWORD_PEPPER string
}

type DBConfig struct {
	DSN string
}

func init() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
}

func Load() Config {
	cfg := Config{
		Addr: ":" + os.Getenv("ADDR"),
		DB: DBConfig{
			DSN: os.Getenv("DSN"),
		},
	}

	return cfg
}

func GetEnv(key EnvKey) string {
	val := os.Getenv(string(key))
	if val == "" {
		panic("missing environment variable: " + string(key))
	}
	return val
}
