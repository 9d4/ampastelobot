package common

import "testing"

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
