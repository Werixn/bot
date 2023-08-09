package commands

import (
	"encoding/json"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) All(inputeMessage *tgbotapi.Message) {
	outputMsgText := "Вот все ваши сеты"
	sets := c.setService.List()
	for _, p := range sets {
		outputMsgText += "set: " + p.Exercises
		outputMsgText += "\n"
	}
}

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsgText := "Here all the exercises: \n\n"

	exercises := c.exerciseService.List()
	for _, p := range exercises {
		outputMsgText += "exercise: " + p.Title
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CommandData{
		Offset: 21,
	})

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(serializedData)),
		),
	)

	c.bot.Send(msg)
}
