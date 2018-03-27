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

	objects, err := fundaObjectsFromSearchResult(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, objects, fundaObjects{
		&fundaObject{
			id:       4151975,
			address:  "Graaf Janlaan 24",
			price:    "€ 595.000",
			url:      parseURL("https://www.funda.nl/4151975"),
			imageURL: parseURL("http://cloud.funda.nl/valentina_media/092/438/006_720x480.jpg"),

			surfaceArea: "105 m² / 243 m²",
			rooms:       "4 kamers",
		},
		&fundaObject{
			id:          4093224,
			address:     "Rapsodie 19",
			price:       "€ 350.000",
			url:         parseURL("https://www.funda.nl/4093224"),
			imageURL:    parseURL("http://cloud.funda.nl/valentina_media/090/656/773_720x480.jpg"),
			surfaceArea: "148 m² / 174 m²",
			rooms:       "6 kamers",
		},
	})
}

func TestFundaSearchURL(t *testing.T) {
	exp := parseURL("https://mobile.funda.io/api/v1/Aanbod/koop/amsterdam/1-dag?page=1&pageSize=10")
	got, err := fundaSearchURL("/amsterdam/1-dag", 1, 10)

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
