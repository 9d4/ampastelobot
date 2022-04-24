package script

import (
	"github.com/9d4/ampastelobot/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Script struct {
	Name string
	Text string
}

func New(name string, text string) *Script {
	return &Script{
		Name: name,
		Text: text,
	}
}

func (s *Script) Run() {

}

func (s *Script) SaveToDB(update tgbotapi.Update) error {
	userID := update.Message.From.ID

	stmt, err := database.DB.Prepare(`INSERT INTO scripts (user_id, name, text) VALUES (?,?,?)`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userID, s.Name, s.Text)
	if err != nil {
		return err
	}

	return nil
}
