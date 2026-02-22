package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvKey string

const (
	EnvAddr   EnvKey = "ADDR"
	EnvDSN    EnvKey = "DSN"
	EnvPepper EnvKey = "PASSWORD_PEPPER"
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
