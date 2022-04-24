package commands

import (
	"github.com/9d4/ampastelobot/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Matcher interface {
	Match()
}

const ScriptHelpRst string = `
<u><b>Script</b></u>

Usage:
<b>Save script</b>
<code>/script save name /req google.com</code>

<b>Run script</b>
<code>/script run name</code> 

Example:
<code>/req https://google.com</code>
<code>/request https://google.com</code>
`

func Script(bot *tgbotapi.BotAPI, update tgbotapi.Update, m Matcher) {
	cmd := common.ParseCommand(update.Message.Text)

	// route the subcommand
	switch cmd.Subcommand {
	case "", "help":
		scriptSendHelp(bot, update)
	}
}

func scriptSendHelp(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, ScriptHelpRst)
	msg.ParseMode = "HTML"

	bot.Send(msg)
}
