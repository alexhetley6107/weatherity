package main

import (
	"log"
	"os"

	"github.com/alexhetley6107/weatherity/internal/weather"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env file not found, using system environment")
	}

	weatherToken := os.Getenv("OPENWEATHER_TOKEN")

	res, err := weather.GetWeather("Лондон", weatherToken)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)

}