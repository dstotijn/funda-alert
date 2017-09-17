package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

var telegramBaseURL = "https://api.telegram.org"

func (object *fundaObject) sendToTelegram(chatID int, botToken string) error {
	data := url.Values{
		"chat_id":    []string{strconv.Itoa(chatID)},
		"parse_mode": []string{"HTML"},
		"text":       []string{object.telegramText()},
	}

	resp, err := http.PostForm(telegramBaseURL+"/bot"+botToken+"/sendMessage", data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("unexpected HTTP response: %s", body)
	}

	return nil
}

func (object *fundaObject) telegramText() string {
	return fmt.Sprintf(`<a href="%v">&#8205;</a><a href="%v">%v</a>
%v kamer(s), %v m²
<strong>%v</strong>`,
		object.imageURL.String(),
		object.url.String(),
		object.address,
		object.numberOfRooms,
		object.surfaceArea,
		object.price,
	)
}
