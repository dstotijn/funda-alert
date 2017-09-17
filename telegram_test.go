package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendTelegramMessages(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"ok":true}`)
	}))
	defer ts.Close()

	telegramBaseURL = ts.URL

	object := &fundaObject{}
	err := object.sendToTelegram(42, "foobar")
	assert.Nil(t, err)
}

func TestSendTelegramMessagesError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer ts.Close()

	telegramBaseURL = ts.URL

	object := &fundaObject{}
	err := object.sendToTelegram(0, "")
	assert.NotNil(t, err, "HTTP response with non `200 OK` should result in an error.")
}

func TestTelegramText(t *testing.T) {
	object := &fundaObject{
		id:            "d113f0dd-4c05-4984-92ca-f7c739623dec",
		address:       "Hoofdweg 99 - C",
		price:         "€ 300.000 k.k.",
		url:           parseURL("http://www.funda.nl/koop/amsterdam/appartement-49397570-hoofdweg-99-c/"),
		imageURL:      parseURL("http://cloud.funda.nl/valentina_media/085/371/511_middel.jpg"),
		surfaceArea:   65,
		numberOfRooms: 3,
	}

	exp := `<a href="http://cloud.funda.nl/valentina_media/085/371/511_middel.jpg">&#8205;</a><a href="http://www.funda.nl/koop/amsterdam/appartement-49397570-hoofdweg-99-c/">Hoofdweg 99 - C</a>
3 kamer(s), 65 m²
<strong>€ 300.000 k.k.</strong>`
	got := object.telegramText()

	assert.Equal(t, exp, got)
}
