package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Addr string
	DB   DBConfig
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
