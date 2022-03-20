package main

import (
	"log"
	"os"
	"sync"

	"github.com/9d4/ampastelobot/common"
	"github.com/9d4/ampastelobot/database"
	"github.com/9d4/ampastelobot/matcher"
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
				matcher := matcher.New(bot, update)
				matcher.Match()
			}()
		}
	}()

	log.Println("Bot ready!")
	wg.Wait()
}
