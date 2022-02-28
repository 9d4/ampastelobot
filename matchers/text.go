package matchers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/traperwaze/ampastelobot/action"
)

func Text(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// not implemented yet
}

func TextAction(botUpdate *action.BotUpdate) bool {
	Text(botUpdate.Bot, *botUpdate.Update)
	return true
}
