package common

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStripBackspaces(t *testing.T) {
	t.Parallel()
	want := "/start ampas tebu"

	for _, test := range []string{
		"/start ampas tebu ",
		"/start ampas    tebu ",
		"/start ampas tebu    ",
		"/start  ampas  tebu  ",
		"  /start ampas  tebu ",
		" /start  ampas  tebu",
		" /start  ampas 	tebu",
		"/start  ampas 	tebu		",
		`/start  ampas    tebu`,
	} {
		result := StripSpaces(test)

		if result != want {
			t.Errorf("got '%v', want '%v'", result, want)
		}
	}
}

func TestParseCommand(t *testing.T) {
	t.Parallel()

	type test struct {
		want  Command
		input string
	}

	for _, item := range []test{
		{want: Command{Command: "start", Subcommand: "ampas", Args: []string{"ampas", "-o", "asdl"}}, input: "/start ampas         ampas -o asdl"},
	} {
		cmd := ParseCommand(item.input)

		if !reflect.DeepEqual(cmd, item.want) {
			t.Errorf("want: '%v', got:'%v'", item.want, cmd)
		}
	}

}

func TestCmdHasArg(t *testing.T) {
	cmd := ParseCommand("/ampas mm -m asdad")

	fmt.Println("jajan", cmd.getArg("asdad"), cmd.GetArgValue("-m"))
}
