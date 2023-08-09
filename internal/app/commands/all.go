package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander) All(inputeMessage *tgbotapi.Message) {
	outputMsgText := "Вот все ваши сеты"
	sets := c.setService.List()
	for _, p := range sets {
		outputMsgText += "set: " + p.Exercises
		outputMsgText += "\n"
	}
}
