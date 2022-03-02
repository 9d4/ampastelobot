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

func (cmd *Command) getArg(arg string) int {
	for i, a := range cmd.Args {
		if a == arg {
			return i
		}
	}

	return -1
}

// get the value of given arg string from command.
// e.g. -m GET
// GetArgValue("-m") will return GET.
// if arg or value not found return empty string
func (cmd *Command) GetArgValue(arg string) (ret string) {
	i := cmd.getArg(arg)
	if i == -1 {
		return
	}

	// value should be in [i+1]
	if len(cmd.Args)-1 <= i {
		return
	}

	ret = cmd.Args[i+1]	

	return
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

// wrapper to send html parsed text message to user easily
func SendMessageHtml(bot *tgbotapi.BotAPI, chatID int64, msg string) {
	mc := tgbotapi.NewMessage(chatID, msg)
	mc.ParseMode = "HTML"
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
