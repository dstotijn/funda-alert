package main

import (
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if _, _, err := fundaObjects(res.Body); err != nil {
		log.Fatal(err)
	}
}
