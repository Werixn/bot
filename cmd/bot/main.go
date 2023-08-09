package main

import (
	"log"
	"os"

	"github.com/werixn/bot/internal/app/commands"
	"github.com/werixn/bot/internal/service/exercise"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	exerciseService := exercise.NewService()

	commander := commands.NewCommander(bot, exerciseService)

	for update := range updates {
		commander.HandleUpdate(update)
	}
}
