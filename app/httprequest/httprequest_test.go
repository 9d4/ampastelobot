package httprequest

import (
	"testing"
)

func TestDoSimple(t *testing.T) {
	_, err := NewSimpleRequest("https://google.com").DoSimple()
	if err != nil {
		t.Error(err)
	}
}

func TestDo(t *testing.T) {
	_, err := NewSimpleRequest("https://google.com").Do()
	if err != nil {
		t.Error(err)
	}
}

func TestCheckUrl(t *testing.T) {
	for _, hr := range []*HttpRequest{
		NewSimpleRequest("google.com"),
		NewSimpleRequest("localhost"),
		NewSimpleRequest("https://google.com"),
		NewSimpleRequest("http://goo.gle"),
		NewSimpleRequest("goo.gle"),
		NewSimpleRequest("goo.gle/wkwkwkwk/asdmlasmd/askdk"),
		NewSimpleRequest("goo.gle:80"),
	}{
		err := hr.checkUrl()

		if err != nil {
			t.Fatal(err)
		}
	}
}
