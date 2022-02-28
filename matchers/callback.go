package matchers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/traperwaze/ampastelobot/action"
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

// CallbackQuery wrapper that supports ActionFunc
func CallbackQueryAction(botUpdate *action.BotUpdate) bool {
	bot, update := botUpdate.Bot, botUpdate.Update

	if update.CallbackQuery != nil {
		CallbackQuery(bot, *update)
		return false
	}

	return true
}
