package matchers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	cmd "github.com/traperwaze/ampastelobot/commands"
	"github.com/traperwaze/ampastelobot/session"
)

func Command(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		log.Println("[session]", session.CreateSession(update).SessionID)
		// bot.Send(handleStart(update))
	case "menu":
		cmd.Menu(bot, update)
	default:
		// bot.Send(handleUnknownCommand(update))
	}
}
