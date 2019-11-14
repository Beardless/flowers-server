package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	Host     string
	User     string
	Password string
}

func ReturnEnvs() EnvVariables {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envs := EnvVariables{
		Host:     os.Getenv("HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("PASSWORD"),
	}

	return envs
}
