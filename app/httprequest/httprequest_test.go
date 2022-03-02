package httprequest

import (
	"testing"
)

func TestDoSimple(t *testing.T) {
	t.Parallel()

	_, err := NewSimpleRequest("google.com").DoSimple()
	if err != nil {
		t.Error(err)
	}
}

func TestDo(t *testing.T) {
	t.Parallel()
	
	_, err := NewSimpleRequest("google.com").Do()
	if err != nil {
		t.Error(err)
	}
}

func TestCheckUrl(t *testing.T) {
	t.Parallel()

	for _, hr := range []*HttpRequest{
		NewSimpleRequest("google.com"),
		NewSimpleRequest("localhost"),
		NewSimpleRequest("https://google.com"),
		NewSimpleRequest("http://goo.gle"),
		NewSimpleRequest("goo.gle"),
		NewSimpleRequest("goo.gle/wkwkwkwk/asdmlasmd/askdk"),
		NewSimpleRequest("goo.gle:80"),
	} {
		err := hr.checkUrl()

		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestDoWithNewRequest(t *testing.T) {
	t.Parallel()

	for _, hr := range []*HttpRequest{
		NewRequest("google.com", "get"),
		NewRequest("google.com", "head"),
		NewRequest("ask.com", "GET"),
		NewRequest("about.google/stories", "GET"),
	}{
		_, err := hr.Do()
		if err != nil {
			t.Fatal(err)
		}
	}
}
