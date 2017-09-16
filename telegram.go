package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

var telegramBaseURL = "https://api.telegram.org"

func (objects fundaObjects) sendTelegramMessages(chatID int, botToken string) error {
	for _, object := range objects {
		data := url.Values{
			"chat_id":    []string{strconv.Itoa(chatID)},
			"parse_mode": []string{"HTML"},
			"text": []string{fmt.Sprintf(`<a href="%v">&#8205;</a><a href="%v">%v</a>
%v`, object.imageURL.String(), object.url.String(), object.address, object.price)},
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
	}

	return nil
}
