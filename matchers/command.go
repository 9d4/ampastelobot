package matchers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/traperwaze/ampastelobot/action"
	cmd "github.com/traperwaze/ampastelobot/commands"
)

func Command(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		cmd.Start(bot, update)
	case "script":
		cmd.Script(bot, update)
	case "req", "request":
		cmd.HttpRequest(bot, update)
	}
}

func CommandAction(botUpdate *action.BotUpdate) bool {
	bot, update := botUpdate.Bot, *botUpdate.Update

	if update.Message.IsCommand() {
		Command(bot, update)

		return false
	}

	return true
}
