package main

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func handleNotFound(update tgbotapi.Update) tgbotapi.Chattable {
	msgConf := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msgConf.Text = "You have typed something in wrong format :("
	msgConf.ReplyToMessageID = update.Message.MessageID

	return msgConf
}

func handleUnknownCommand(update tgbotapi.Update) tgbotapi.Chattable {
	msgConf := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msgConf.Text = "There's no such command 😕"
	msgConf.ReplyToMessageID = update.Message.MessageID

	return msgConf
}

func handleStart(update tgbotapi.Update) tgbotapi.Chattable {
	msgConf := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msgConf.Text = "Bot started! type /menu to get menu list"
	msgConf.ReplyToMessageID = update.Message.MessageID

	return msgConf
}

func handleCommandMenu(update tgbotapi.Update) tgbotapi.Chattable {
	msgConf := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msgConf.Text = "Main Menu"

	return msgConf
}

func handleMenuBlynk(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

}
