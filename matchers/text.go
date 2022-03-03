package matchers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/9d4/ampastelobot/action"
)

func Text(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// not implemented yet
}

func TextAction(botUpdate *action.BotUpdate) bool {
	Text(botUpdate.Bot, *botUpdate.Update)
	return true
}
