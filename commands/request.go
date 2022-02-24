package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/traperwaze/ampastelobot/app/httprequest"
	"github.com/traperwaze/ampastelobot/common"
)

// use ParseMode = "HTML"
const HttpRequestHelpRst string = `
<u><b>Request</b></u>

Make simple head request to a url or address.

Usage:
/req URL

Example:
<code>/req https://google.com</code>
<code>/request https://google.com</code>
`

func HttpRequest(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	cmd := common.ParseCommand(update.Message.Text)

	// we assume the cmd.Subcommand as the url of the request
	// e.g /req 	https://google.com
	//     cmd		url
	// cmd. Cmd		Subcommand

	fallback := func() {
		HttpRequestSendHelp(bot, update)
	}

	if cmd.Subcommand == "" {
		fallback()
		return
	}

	if len(cmd.Args) == 0 {
		// do simple request
		body, err := httprequest.NewSimpleRequest(cmd.Subcommand).DoSimple()
		if err != nil {
			common.SendMessageText(bot, update.Message.Chat.ID, "Couldn't make request")
			HttpRequestSendHelp(bot, update)
			return
		}

		common.SendMessageText(bot, update.Message.Chat.ID, fmt.Sprint(body))
		return
	}
}

func HttpRequestSendHelp(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	help := tgbotapi.NewMessage(update.Message.Chat.ID, HttpRequestHelpRst)
	help.ParseMode = "HTML"

	bot.Send(help)
}
