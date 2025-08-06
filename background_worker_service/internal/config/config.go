package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	CurrencyAPIURL string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Файл .env не найден")
	}

	return &Config{
		CurrencyAPIURL: os.Getenv("CURRENCY_API_URL"),
	}
}
