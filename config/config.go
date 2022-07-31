package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var C *schema

type schema struct {
	Db_url string
	Port   string
}

func getEnv(key string) string {
	value := os.Getenv(key)
	return value
}

func setupEnv() (string, string, error) {
	var err error = nil
	if getEnv("APP_ENV") != "PRODUCTION" {
		err = godotenv.Load(".env")
	}

	if err != nil {
		log.Error().Msg("Error loading .env file")
		return "", "", err
	}

	url := fmt.Sprintf("host=%v port=%v user=%v "+
		"password=%v dbname=%v sslmode=disable",
		getEnv("POSTGRES_HOST"), getEnv("POSTGRES_PORT"), getEnv("POSTGRES_USER"), getEnv("POSTGRES_PASSWORD"), getEnv("POSTGRES_DB"))

	port := "3001"
	return url, port, nil
}

func Init() {
	db_url, port, err := setupEnv()
	if err != nil {
		log.Error().Msg("Error initializing db")
		panic(err)
	}

	C = &schema{
		Db_url: db_url,
		Port:   port,
	}
}
