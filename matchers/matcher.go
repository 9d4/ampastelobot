package matchers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Match(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.CallbackQuery != nil {
		CallbackQuery(bot, update)
		return
	}

	if update.Message == nil {
		return
	}

	if update.Message.IsCommand() {
		Command(bot, update)
		return
	}

	Text(bot, update)
}
