package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/9d4/ampastelobot/database"
	"github.com/9d4/ampastelobot/matchers"
)

func init() {
	godotenv.Load()
	database.Init()
}

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// set debug mode
	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	go func() {
		for {
			update := <-updates

			go func() { matchers.Match(bot, update) }()
		}
	}()

	log.Println("Bot ready!")
	select {}
}
