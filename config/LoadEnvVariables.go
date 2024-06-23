package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() error {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
		return err
	}
	return nil
}
