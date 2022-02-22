package common

import (
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

	type ret struct {
		cmd  string
		args []string
	}

	type test struct {
		want  ret
		input string
	}

	for _, item := range []test{
		{want: ret{cmd: "start", args: []string{"ampas", "ampas", "-o", "asdl"}}, input: "/start ampas ampas -o asdl"},
		{want: ret{cmd: "start", args: []string{"save", "-0", "things"}}, input: "/start save -0 things"},
		{want: ret{cmd: "start", args: []string{"kuda", "antariksa"}}, input: "/start kuda   antariksa   "},
		{want: ret{cmd: "start", args: []string{"uyeeeeee", "okokok"}}, input: "   /start uyeeeeee   okokok"},
	} {
		cmd, args := ParseCommand(item.input)

		if !(cmd == item.want.cmd && reflect.DeepEqual(args, item.want.args)) {
			tmp := ret{cmd: cmd, args: args}

			t.Errorf("want: '%v', got:'%v'", item.want, tmp)
		}
	}

}
