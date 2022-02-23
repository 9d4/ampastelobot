package common

import (
	"os"
	"strings"
	"unicode"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Command struct {
	Command    string
	Subcommand string
	Args       []string
}

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

// returning the command and it's args
// e.g. /script ampas tebu
// returns "script" ["ampas", "tebu"]
// 			string	[]string
func ParseCommand(cmd string) (command Command) {
	stripped := StripSpaces(cmd)

	// split string by space
	slice := strings.Split(stripped, " ")

	// assign the command (slice[0]) to command.Command
	// but remove the slash
	command.Command = string([]byte(slice[0])[1:])

	if len(slice) >= 2 {
		// sub command is index 1 of slice
		command.Subcommand = slice[1]
	}

	if len(slice) >= 3 {
		// args starts from index 2 of slice
		command.Args = slice[2:]
	}

	return
}

func StripSpaces(text string) string {
	trimmed := []byte(strings.TrimSpace(text))
	out := []byte{}

	// then trim duplicate space in the center of the text
	gotSpace := false
	for _, l := range trimmed {
		if !unicode.IsSpace(rune(l)) {
			out = append(out, l)

			gotSpace = false
		}

		if unicode.IsSpace(rune(l)) {
			// if previously got space, don't include the next space
			if !gotSpace {
				out = append(out, l)
			}

			gotSpace = true
		}
	}

	return string(out)
}
