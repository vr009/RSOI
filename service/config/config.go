package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	godotenv.Load(".env")
}

func GetConnectionString() (string, error) {
	os.Setenv("DATABASE_URL", "dbname=persons host=127.0.0.1 user=slavaryanov password=postgres port=5432 pool_max_conns=10")
	key, flag := os.LookupEnv("DATABASE_URL")
	if !flag {
		return "", errors.New("connection string not found")
	}
	return key, nil
}
