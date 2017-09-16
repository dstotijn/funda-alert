package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	telegramChatID := flag.Int("telegramChatID", 0, "Telegram \"chat_id\" to send messages to")
	fundaSearchOptions := flag.String("fundaSearchOptions", "/amsterdam/1-dag", "Funda search options")
	flag.Parse()

	if *telegramChatID == 0 {
		log.Fatal("Error: Missing parameter `telegramChatID`.")
	}

	fundaToken := os.Getenv("FUNDA_ALERT_FUNDA_TOKEN")
	if fundaToken == "" {
		log.Fatal("Error: Environment variable `FUNDA_ALERT_FUNDA_TOKEN` cannot be empty.")
	}

	telegramToken := os.Getenv("FUNDA_ALERT_TELEGRAM_TOKEN")
	if telegramToken == "" {
		log.Fatal("Error: Environment variable `FUNDA_ALERT_TELEGRAM_TOKEN` cannot be empty.")
	}

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		log.Fatal(err)
	}

	u, err := fundaSearchURL(fundaToken, *fundaSearchOptions, 1, 10)
	if err != nil {
		log.Fatal(err)
	}

	req.URL = u

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: Unexpected HTTP response code `%d` received.", resp.StatusCode)
	}

	objects, _, err := fundaObjectsFromSearchResult(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Retrieved %d object(s).", len(objects))

	if err := objects.sendTelegramMessages(*telegramChatID, telegramToken); err != nil {
		log.Fatal(err)
	}
}
