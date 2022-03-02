package commands

import (
	"fmt"
	"strings"

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

	hr := httprequest.NewSimpleRequest(cmd.Subcommand)

	// if there's no args supplied then
	// do basic simple request where it uses HEAD method
	if len(cmd.Args) == 0 {
		body, err := hr.DoSimple()
		if err != nil {
			common.SendMessageText(bot, update.Message.Chat.ID, "Couldn't make request")
			HttpRequestSendHelp(bot, update)
			return
		}

		common.SendMessageText(bot, update.Message.Chat.ID, fmt.Sprint(body))
		return
	}

	if method := cmd.GetArgValue("-m"); method != "" {
		hr.Method = strings.ToUpper(method)

		statusCode, err := hr.DoSimple()
		if err != nil {
			common.SendMessageText(bot, update.Message.Chat.ID, "Couldn't make request")
			HttpRequestSendHelp(bot, update)
			return
		}

		common.SendMessageText(bot, update.Message.Chat.ID, fmt.Sprint(statusCode))
		return
	}
}

// send help message to user about:
// /req, /request
func HttpRequestSendHelp(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	help := tgbotapi.NewMessage(update.Message.Chat.ID, HttpRequestHelpRst)
	help.ParseMode = "HTML"

	bot.Send(help)
}
