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

	// actuator.Add(matchers.CallbackQueryAction)
	actuator.Add(m.CommandAction)
	// actuator.Add(matchers.TextAction)

	actuator.Exec()
}

func (m *Matcher) MatchNew(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	matcher := New(bot, update)
	matcher.Match()
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
