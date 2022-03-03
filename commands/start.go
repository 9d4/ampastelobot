package commands

import (
	"reflect"

	"github.com/9d4/ampastelobot/action"
	"github.com/9d4/ampastelobot/common"
)

func Start(botUpdate *action.BotUpdate) {
	bot, update := botUpdate.Bot, *botUpdate.Update

	// check first_time or not
	first_time := false

	if botUpdate.Data["first_time"] != nil &&
		reflect.ValueOf(botUpdate.Data["first_time"]).Kind() == reflect.Bool &&
		botUpdate.Data["first_time"] == true {
		first_time = true
	}

	if !first_time {
		common.SendMessageText(bot, update.Message.Chat.ID, "We already have session")
		return
	}

	common.SendMessageText(bot, update.Message.Chat.ID, "Hello the internet")
}
