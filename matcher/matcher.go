package matcher

import (
	"github.com/9d4/ampastelobot/action"
	cmd "github.com/9d4/ampastelobot/commands"
	"github.com/9d4/ampastelobot/common"
	"github.com/9d4/ampastelobot/session"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Matcher struct {
	bot    *tgbotapi.BotAPI
	update tgbotapi.Update
}

func New(bot *tgbotapi.BotAPI, update tgbotapi.Update) *Matcher {
	return &Matcher{
		bot:    bot,
		update: update,
	}
}

func (m *Matcher) Match() {
	actuator := action.NewActuator(m.bot, m.update)

	actuator.Add(session.Middleware)

	actuator.Add(m.CallbackQueryAction)
	actuator.Add(m.CommandAction)
	actuator.Add(m.TextAction)

	actuator.Exec()
}

func (m *Matcher) Command(botUpdate *action.BotUpdate) {
	bot, update := botUpdate.Bot, *botUpdate.Update

	command := common.ParseCommand(update.Message.Text)

	switch command.Command {
	case "start":
		cmd.Start(botUpdate)
	case "script":
		cmd.Script(bot, update, m)
	case "req", "request":
		cmd.HttpRequest(bot, update)
	}
}

func (m *Matcher) CommandAction(botUpdate *action.BotUpdate) bool {
	_, update := botUpdate.Bot, *botUpdate.Update

	if update.Message.IsCommand() {
		m.Command(botUpdate)

		return false
	}

	return true
}

func (m *Matcher) CallbackQuery(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	cbQuery := update.CallbackQuery
	cbData := cbQuery.Data

	// to make sure bot not hanging, send ack to the bot
	bot.Request(tgbotapi.NewCallback(cbQuery.ID, cbData))

	// to make everything clear, delete the callback message from bot
	delMsgConf := tgbotapi.NewDeleteMessage(cbQuery.Message.Chat.ID, cbQuery.Message.MessageID)
	bot.Request(delMsgConf)

	// then do the rest
}

// CallbackQuery wrapper that supports ActionFunc
func (m *Matcher) CallbackQueryAction(botUpdate *action.BotUpdate) bool {
	bot, update := botUpdate.Bot, botUpdate.Update

	if update.CallbackQuery != nil {
		m.CallbackQuery(bot, *update)
		return false
	}

	return true
}

func (m *Matcher) Text(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// not implemented yet
}

func (m *Matcher) TextAction(botUpdate *action.BotUpdate) bool {
	m.Text(botUpdate.Bot, *botUpdate.Update)
	return true
}
