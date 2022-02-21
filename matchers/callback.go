package matchers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CallbackQuery(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	cbQuery := update.CallbackQuery
	cbData := cbQuery.Data

	// to make sure bot not hanging, send ack to the bot
	bot.Request(tgbotapi.NewCallback(cbQuery.ID, cbData))

	// to make everything clear, delete the callback message from bot
	delMsgConf := tgbotapi.NewDeleteMessage(cbQuery.Message.Chat.ID, cbQuery.Message.MessageID)
	bot.Request(delMsgConf)

	// then do the rest
}
