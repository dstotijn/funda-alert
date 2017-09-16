package main

import (
	"flag"
	"log"
	"os"

	"github.com/boltdb/bolt"
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

	r, err := searchFunda(fundaToken, *fundaSearchOptions, 1, 25)
	if err != nil {
		log.Fatalf("Error: Could not search Funda: %v", err)
	}
	defer r.Close()

	objects, _, err := fundaObjectsFromSearchResult(r)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Retrieved %d object(s).", len(objects))

	for _, object := range objects {
		err := db.Update(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte(dbBucket))

			id := bucket.Get([]byte(object.id))
			if id != nil {
				log.Printf("Skipping object (%v), already handled.", object.id)
				return nil
			}

			if err := object.sendToTelegram(*telegramChatID, telegramToken); err != nil {
				return err
			}
			log.Printf("Sent message for object (%v) to Telegram.", object.id)

			return bucket.Put([]byte(object.id), nil)
		})
		if err != nil {
			log.Fatalf("Error: Could not handle Funda object: %v", err)
		}
	}
}
