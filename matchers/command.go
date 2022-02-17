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
		if _, err :=session.CreateSession(update); err == nil {
			log.Println("[command] succesfully creating session")
		}

		// bot.Send(handleStart(update))
	case "menu":
		cmd.Menu(bot, update)
	default:
		// bot.Send(handleUnknownCommand(update))
	}
}
