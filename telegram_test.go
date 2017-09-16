package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendTelegramMessages(t *testing.T) {
	objects := fundaObjects{
		&fundaObject{
			address:  "Jan Evertsenstraat 137 F 1057 BV Amsterdam",
			price:    "€ 250.000 k.k.",
			url:      parseURL("https://www.funda.nl/koop/amsterdam/appartement-49382543-jan-evertsenstraat-137-f/"),
			imageURL: parseURL("https://cloud.funda.nl/valentina_media/085/222/384_360x240.jpg"),
		},
		&fundaObject{
			address:  "Tweede Atjehstraat 26 II 1094 LG Amsterdam",
			price:    "€ 325.000 k.k.",
			url:      parseURL("https://www.funda.nl/koop/amsterdam/appartement-49389196-tweede-atjehstraat-26-ii/"),
			imageURL: parseURL("https://cloud.funda.nl/valentina_media/085/236/192_360x240.jpg"),
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"ok":true}`)
	}))
	defer ts.Close()

	telegramBaseURL = ts.URL

	err := objects.sendTelegramMessages(42, "foobar")
	assert.Nil(t, err)
}
