package bot

import (
	"fmt"
	"log"

	"github.com/alexhetley6107/weatherity/internal/env"
	"github.com/alexhetley6107/weatherity/internal/weather"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api    *tgbotapi.BotAPI
	config *env.Config
}

func NewBot(cfg *env.Config) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		return nil, err
	}

	return &Bot{api: api, config: cfg}, nil
}

func (b *Bot) Start() {
	log.Printf("Authorized on account %s", b.api.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := b.api.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		text := update.Message.Text

		switch text {
		case "/start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Напиши город — я покажу погоду")
			b.api.Send(msg)
		default:
			result, err := weather.GetWeather(text, b.config.OpenWeatherToken)
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Ошибка: %v", err))
				b.api.Send(msg)
				continue
			}
			b.api.Send(tgbotapi.NewMessage(update.Message.Chat.ID, result))
		}
	}
}
