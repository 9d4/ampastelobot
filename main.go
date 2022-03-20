package main

import (
	"log"
	"os"
	"sync"

	"github.com/9d4/ampastelobot/action"
	"github.com/9d4/ampastelobot/common"
	"github.com/9d4/ampastelobot/database"
	"github.com/9d4/ampastelobot/matchers"
	"github.com/9d4/ampastelobot/session"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	database.Init()
}

func main() {
	var wg sync.WaitGroup

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// set debug mode
	if common.IsDevelopment() {
		bot.Debug = true
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			update := <-updates

			go func() {
				actuator := action.NewActuator(bot, update)

				actuator.Add(session.Middleware)

				actuator.Add(matchers.CallbackQueryAction)
				actuator.Add(matchers.CommandAction)
				actuator.Add(matchers.TextAction)

				actuator.Exec()
			}()
		}
	}()

	log.Println("Bot ready!")
	wg.Wait()
}
