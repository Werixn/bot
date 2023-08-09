package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/werixn/bot/internal/service/exercise"
	"github.com/werixn/bot/internal/service/set"
)

type Commander struct {
	bot             *tgbotapi.BotAPI
	exerciseService *exercise.Service
}

type Commander2 struct {
	bot        *tgbotapi.BotAPI
	setService *set.Service
}

func NewSecondCommander(
	bot *tgbotapi.BotAPI,
	setService *set.Service,
) *Commander2 {
	return &Commander2{
		bot:        bot,
		setService: setService,
	}
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
