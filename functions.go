package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func handleNotFound(update tgbotapi.Update) tgbotapi.Chattable {
	msgConf := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msgConf.Text = "You have typed something in wrong format :("
	msgConf.ReplyToMessageID = update.Message.MessageID

	return msgConf
}
