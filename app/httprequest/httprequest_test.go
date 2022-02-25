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
