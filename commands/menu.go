package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// menu.go
// this exclusively handles "/menu" command and it's childs

var MainMenuMarkup tgbotapi.InlineKeyboardMarkup = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Blynk", "menu_blynk"),
		tgbotapi.NewInlineKeyboardButtonData("OOT", "menu_oot"),
	),
)

func Menu(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

}
