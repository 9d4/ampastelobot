package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var MainMenuMarkup tgbotapi.InlineKeyboardMarkup = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		// tgbotapi.NewInlineKeyboardButtonData("", ""),
	),
)

func Menu(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msgConf := tgbotapi.NewMessage(update.Message.Chat.ID, "Menu")
	msgConf.ReplyMarkup = MainMenuMarkup

	bot.Send(msgConf)
}
