package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
		return "", err
	}
	return os.Getenv(key), nil
}