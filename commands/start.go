package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/traperwaze/ampastelobot/common"
	"github.com/traperwaze/ampastelobot/session"
)

func Start(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// check wheter has session or not
	sess, err := session.GetSession(update)
	if err != nil {
		switch err {
		case err.(*session.ErrNoSessionInDB):
			// if error is ErrNoSessionInDB
			// then create session
			sess, _ = session.CreateSession(update)
		default:
			log.Println(err)
			return
		}
	}

	// more good stuff
	if err == nil && sess.UserID > 0 {
		common.SendMessageText(bot, update.Message.Chat.ID, "We already have session")
		return
	}

	common.SendMessageText(bot, update.Message.Chat.ID, "Hello the internet")
}
