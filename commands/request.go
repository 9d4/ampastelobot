package commands

import (
	"fmt"

	"github.com/9d4/ampastelobot/app/httprequest"
	"github.com/9d4/ampastelobot/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

	if cmd.Subcommand == "" {
		HttpRequestSendHelp(bot, update)
		return
	}

	hr := httprequest.NewSimpleRequest(cmd.Subcommand)
	method := cmd.GetArgValue("-m")

	var (
		body   string
		reqErr error
	)

	switch method {
	case "":
		hr.Method = "HEAD"
		statusCode, err := hr.DoSimple()
		// we need to do this due to type incompatibility
		reqErr = err
		body = fmt.Sprint(statusCode)

	default:
		hr.Method = method
		body, reqErr = hr.Do()
	}

	if reqErr != nil {
		common.SendMessageText(bot, update.Message.Chat.ID, "Couldn't make request")
		common.SendMessageText(bot, update.Message.Chat.ID, reqErr.Error())
		return
	}

	common.SendMessageText(bot, update.Message.Chat.ID, body)
}

// send help message to user about:
// /req, /request
func HttpRequestSendHelp(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	help := tgbotapi.NewMessage(update.Message.Chat.ID, HttpRequestHelpRst)
	help.ParseMode = "HTML"

	bot.Send(help)
}
