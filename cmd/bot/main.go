package main

import (
	"log"

	"github.com/alexhetley6107/weatherity/internal/bot"
	"github.com/alexhetley6107/weatherity/internal/env"
)

func main() {

	cfg, err := env.Load()
	if err != nil {
		log.Fatal(err)
	}

	bot, err := bot.NewBot(cfg)
	if err != nil {
		log.Fatal(err)
	}

	bot.Start()

}
