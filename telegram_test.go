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
		id:          4151975,
		address:     "Hoofdweg 99 - C",
		price:       "€ 300.000 k.k.",
		url:         parseURL("https://www.funda.nl/4151975"),
		imageURL:    parseURL("http://cloud.funda.nl/valentina_media/085/371/511_middel.jpg"),
		surfaceArea: "105 m² / 243 m²",
		rooms:       "4 kamers",
	}

	exp := `<a href="http://cloud.funda.nl/valentina_media/085/371/511_middel.jpg">&#8205;</a><a href="https://www.funda.nl/4151975">Hoofdweg 99 - C</a>
4 kamers, 105 m² / 243 m²
<strong>€ 300.000 k.k.</strong>`
	got := object.telegramText()

	assert.Equal(t, exp, got)
}
