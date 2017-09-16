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
