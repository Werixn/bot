package commands

import (
	"github.com/Werixn/bot/internal/service/exercise"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Commander struct {
	bot             *tgbotapi.BotAPI
	exerciseService *exercise.Service
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	exerciseService *exercise.Service,
) *Commander {
	return &Commander{
		bot:             bot,
		exerciseService: exerciseService,
	}
}
