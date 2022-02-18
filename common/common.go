package common

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Wd() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return dir
}

// wrapper to send text message to user easily
func SendMessageText(bot *tgbotapi.BotAPI, chatID int64, msg string) {
	mc := tgbotapi.NewMessage(chatID, msg)
	bot.Send(mc)
}
