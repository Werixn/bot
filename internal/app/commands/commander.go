package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/werixn/bot/internal/service/exercise"
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
