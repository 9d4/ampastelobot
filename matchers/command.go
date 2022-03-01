package matchers

import (
	"github.com/traperwaze/ampastelobot/action"
	cmd "github.com/traperwaze/ampastelobot/commands"
)

func Command(botUpdate *action.BotUpdate) {
	bot, update := botUpdate.Bot, *botUpdate.Update

	switch update.Message.Command() {
	case "start":
		cmd.Start(botUpdate)
	case "script":
		cmd.Script(bot, update)
	case "req", "request":
		cmd.HttpRequest(bot, update)
	}
}

func CommandAction(botUpdate *action.BotUpdate) bool {
	_, update := botUpdate.Bot, *botUpdate.Update

	if update.Message.IsCommand() {
		Command(botUpdate)

		return false
	}

	return true
}
