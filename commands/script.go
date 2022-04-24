package commands

import (
	"strings"

	"github.com/9d4/ampastelobot/app/script"
	"github.com/9d4/ampastelobot/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// used to call the matcher.Match and avoiding import cycle
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
	case "save":
		if len(cmd.Args) < 2 {
			scriptSendHelp(bot, update)
			return
		}

		textCommand := strings.Join(cmd.Args[1:], " ")
		s := script.New(cmd.Args[0], textCommand)
		err := s.SaveToDB(update)
		if err != nil {
			// send "Can't save reason: reason"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
			bot.Send(msg)
		}
	}
}

func scriptSendHelp(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, ScriptHelpRst)
	msg.ParseMode = "HTML"

	bot.Send(msg)
}
