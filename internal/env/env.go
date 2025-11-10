package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken    string
	OpenWeatherToken string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system environment")
	}

	cfg := &Config{
		TelegramToken:    os.Getenv("TELEGRAM_TOKEN"),
		OpenWeatherToken: os.Getenv("OPENWEATHER_TOKEN"),
	}

	if cfg.TelegramToken == "" {
		return nil, fmt.Errorf("TELEGRAM_TOKEN is not set")
	}
	if cfg.OpenWeatherToken == "" {
		return nil, fmt.Errorf("OPENWEATHER_TOKEN is not set")
	}

	return cfg, nil
}
