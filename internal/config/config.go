package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// добавляем элементы env файла в структуру
type Config struct {
	DB_PASSWORD string
	DB_USERNAME string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_SSLMODE  string
}

func ConfigData() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_SSLMODE:  os.Getenv("DB_SSLMODE"),
	}
}
