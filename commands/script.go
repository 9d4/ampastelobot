package commands

import (
	"github.com/9d4/ampastelobot/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Matcher interface {
	MatchNew(*tgbotapi.BotAPI, tgbotapi.Update)
}

func Script(bot *tgbotapi.BotAPI, update tgbotapi.Update, m Matcher) {
	cmd := common.ParseCommand(update.Message.Text)

	println(update.Message.Text)

	update.Message.Text = "/req google.com -m get"

	m.MatchNew(bot, update)

	_ = cmd

	// route the subcommand
	// switch cmd.Subcommand {
	// case "", "list":
	// 	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "HOLLA LIST"))
	// }
}
