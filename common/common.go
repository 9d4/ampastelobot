package common

import (
	"os"
	"strings"
	"unicode"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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
func ParseCommand(cmd string) (command string, args []string) {
	stripped := StripSpaces(cmd)

	// split string by space
	s := strings.Split(stripped, " ")
	command = string([]byte(s[0])[1:]) // remove the slash

	if len(s) > 1 {
		args = s[1:]
		return
	}

	return command, []string{}
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
