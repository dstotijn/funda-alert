package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	telegramChatID := flag.Int("telegramChatID", 0, "Telegram `chat_id` to send messages to")
	flag.Parse()

	if *telegramChatID == 0 {
		log.Fatal("Missing parameter `telegramChatID`")
	}

	resp, err := http.Get("http://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Printf("Received response with HTTP status code: %d", resp.StatusCode)

	objects, _, err := fundaObjectsFromSearchResult(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Retrieved %d object(s).", len(objects))

	if err := objects.sendTelegramMessages(*telegramChatID, os.Getenv("FUNDA_ALERT_TELEGRAM_TOKEN")); err != nil {
		log.Fatal(err)
	}
}
