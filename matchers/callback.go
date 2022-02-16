package matchers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	// cmd "github.com/traperwaze/ampastelobot/commands"
)

func CallbackQuery(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	cbQuery := update.CallbackQuery
	cbData := cbQuery.Data

	bot.Send(tgbotapi.NewMessage(cbQuery.Message.Chat.ID, cbData))
}
