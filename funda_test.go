package main

import (
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFundaObjectsFromSearchResult(t *testing.T) {
	file, err := os.Open("test_data/funda_search_response.json")
	if err != nil {
		t.Fatal(err)
	}

	objects, pageCount, err := fundaObjectsFromSearchResult(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, objects, fundaObjects{
		&fundaObject{
			address:  "Hoofdweg 99 - C",
			price:    "€ 300.000 k.k.",
			url:      parseURL("http://www.funda.nl/koop/amsterdam/appartement-49397570-hoofdweg-99-c/"),
			imageURL: parseURL("http://cloud.funda.nl/valentina_media/085/371/511_middel.jpg"),
		},
		&fundaObject{
			address:  "Geuzenstraat 77 III",
			price:    "€ 250.000 k.k.",
			url:      parseURL("http://www.funda.nl/koop/amsterdam/appartement-49397476-geuzenstraat-77-iii/"),
			imageURL: parseURL("http://cloud.funda.nl/valentina_media/085/368/478_middel.jpg"),
		},
	})

	assert.Equal(t, pageCount, 4)
}

func parseURL(s string) url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	return *u
}
