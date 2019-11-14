package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvVariables struct {
	Host   string
	Port   int
	User   string
	Dbname string
}

func ReturnEnvs() EnvVariables {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	parsedPort, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Couldnt parse port number")
	}

	envs := EnvVariables{
		Host:   os.Getenv("HOST"),
		Port:   parsedPort,
		User:   os.Getenv("USER"),
		Dbname: os.Getenv("DBNAME"),
	}

	return envs
}
