package action

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotUpdate struct {
	Bot    *tgbotapi.BotAPI
	Update *tgbotapi.Update
	Data   map[string]interface{}
}

// ActionFunc is function that receives botAPI and update then return another of them.
// If returned value is true, the Actuator will continue the chain, otherwise not.
type ActionFunc func(*BotUpdate) bool

func NewActuator(botAPI *tgbotapi.BotAPI, update tgbotapi.Update) *Actuator {
	return &Actuator{
		botUpdate: &BotUpdate{
			Bot:    botAPI,
			Update: &update,
			Data:   map[string]interface{}{},
		},
	}
}

type Actuator struct {
	botUpdate *BotUpdate
	actions   []ActionFunc
}

func (atr *Actuator) Add(fn ...ActionFunc) *Actuator {
	atr.actions = append(atr.actions, fn...)
	return atr
}

func (atr *Actuator) Exec() {
	for _, fn := range atr.actions {
		if next := fn(atr.botUpdate); !next {
			break
		}
	}
}
