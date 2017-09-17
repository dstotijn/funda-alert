package main

import (
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			id:            "d113f0dd-4c05-4984-92ca-f7c739623dec",
			address:       "Hoofdweg 99 - C",
			price:         "€ 300.000 k.k.",
			url:           parseURL("http://www.funda.nl/koop/amsterdam/appartement-49397570-hoofdweg-99-c/"),
			imageURL:      parseURL("http://cloud.funda.nl/valentina_media/085/371/511_grotere.jpg"),
			surfaceArea:   65,
			numberOfRooms: 3,
		},
		&fundaObject{
			id:            "27097d93-3547-47a3-ac9f-b16e1112e9ad",
			address:       "Geuzenstraat 77 III",
			price:         "€ 250.000 k.k.",
			url:           parseURL("http://www.funda.nl/koop/amsterdam/appartement-49397476-geuzenstraat-77-iii/"),
			imageURL:      parseURL("http://cloud.funda.nl/valentina_media/085/368/478_grotere.jpg"),
			surfaceArea:   48,
			numberOfRooms: 3,
		},
	})

	assert.Equal(t, pageCount, 4)
}

func TestFundaSearchURL(t *testing.T) {
	exp := parseURL("http://partnerapi.funda.nl/feeds/Aanbod.svc/search/json/foobar/?page=1&pagesize=10&type=koop&website=funda&zo=%2Famsterdam%2F1-dag")
	got, err := fundaSearchURL("foobar", "/amsterdam/1-dag", 1, 10)

	require.Nil(t, err)

	assert.Equal(t, exp, *got)
}

func parseURL(s string) url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	return *u
}
