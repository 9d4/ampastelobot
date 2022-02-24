package httprequest

import (
	"log"
	"testing"
)

func TestDoSimple(t *testing.T) {
	statusCode, err := NewSimpleRequest("https://google.com").DoSimple()
	if err != nil {
		t.Error(err)
	}

	res, err := NewSimpleRequest("https://google.com").Do()
	if err != nil {
		t.Error(err)
	}

	log.Println("Status Code:", statusCode)
	log.Println("Response:", res)
}
