package matchers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/traperwaze/ampastelobot/action"
	"github.com/traperwaze/ampastelobot/session"
)

func Match(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	actuator := action.NewActuator(bot, update)

	actuator.Add(session.Middleware)

	actuator.Add(CallbackQueryAction)
	actuator.Add(CommandAction)
	actuator.Add(TextAction)

	actuator.Exec()
}
