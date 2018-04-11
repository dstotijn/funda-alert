package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
	"github.com/dstotijn/go-funda"
)

const dbBucket = "FundaObjects"

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

	fundaClient := funda.NewClient(fundaToken)

	db, err := bolt.Open("funda_alert.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(dbBucket))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	houses, err := fundaClient.Search(*fundaSearchOptions, 1, 25)
	if err != nil {
		log.Fatalf("Error: Could not search Funda: %v", err)
	}

	log.Printf("Retrieved %d object(s).", len(houses))

	for _, house := range houses {
		err := db.Update(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte(dbBucket))
			key := []byte(fmt.Sprintf("%v:%v", *telegramChatID, house.ID))

			id := bucket.Get(key)
			if id != nil {
				log.Printf("Skipping object (%v), already handled.", house.ID)
				return nil
			}

			if err := sendToTelegram(*house, *telegramChatID, telegramToken); err != nil {
				return err
			}
			log.Printf("Sent message for object (%v) to Telegram.", house.ID)

			return bucket.Put(key, nil)
		})
		if err != nil {
			log.Fatalf("Error: Could not handle Funda object: %v", err)
		}
	}
}
