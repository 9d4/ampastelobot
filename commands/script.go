package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/9d4/ampastelobot/common"
)

func Script(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	cmd := common.ParseCommand(update.Message.Text)

	// route the subcommand
	switch cmd.Subcommand {
	case "", "list":
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "HOLLA LIST"))
	}
}
