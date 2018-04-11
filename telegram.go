package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/dstotijn/go-funda"
)

var telegramBaseURL = "https://api.telegram.org"

func sendToTelegram(house funda.House, chatID int, botToken string) error {
	data := url.Values{
		"chat_id":    []string{strconv.Itoa(chatID)},
		"parse_mode": []string{"HTML"},
		"text":       []string{telegramText(house)},
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

func telegramText(house funda.House) string {
	return fmt.Sprintf(`<a href="%v">&#8205;</a><a href="%v">%v</a>
%v, %v
<strong>%v</strong>`,
		house.ImageURL.String(),
		house.URL.String(),
		house.Address,
		house.Rooms,
		house.SurfaceArea,
		house.Price,
	)
}
