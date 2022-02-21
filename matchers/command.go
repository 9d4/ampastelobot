package matchers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	cmd "github.com/traperwaze/ampastelobot/commands"
)

func Command(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		cmd.Start(bot, update)
	case "script":
		cmd.Script(bot, update)
	}
}
