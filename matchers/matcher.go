package matchers

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/traperwaze/ampastelobot/action"
)

func Match(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	test := func(up *action.BotUpdate) bool {
		fmt.Println("SUCCESS")
		return true
	}


	action.NewActuator(bot, update).Add(test, CallbackQueryAction).Exec()

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
