package matchers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/traperwaze/ampastelobot/action"
)

func Match(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	actuator := action.NewActuator(bot, update)

	actuator.Add(CallbackQueryAction)
	actuator.Add(CommandAction)

	actuator.Exec()

	// atr.Print()

	// if update.CallbackQuery != nil {
	// 	CallbackQuery(bot, update)
	// 	return
	// }

	// if update.Message == nil {
	// 	return
	// }

	// if update.Message.IsCommand() {
	// 	Command(bot, update)
	// 	return
	// }

	// Text(bot, update)
}
