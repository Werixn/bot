package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (c *Commander2) All(inputMessage *tgbotapi.Message) {
	outputMsgText := "Вот все ваши сеты"
	sets := c.setService.All()
	for _, p := range sets {
		outputMsgText += p.Exercises
		outputMsgText += "\n"
	}
}
