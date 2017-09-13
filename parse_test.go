package main

import (
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseFundaObjects(t *testing.T) {
	file, err := os.Open("data/funda.html")
	if err != nil {
		t.Fatal(err)
	}

	objects, nextURL, err := fundaObjects(file)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, objects, []*fundaObject{
		&fundaObject{
			address:  "Jan Evertsenstraat 137 F 1057 BV Amsterdam",
			price:    "€ 250.000 k.k.",
			url:      parseURL("https://www.funda.nl/koop/amsterdam/appartement-49382543-jan-evertsenstraat-137-f/"),
			imageURL: parseURL("https://cloud.funda.nl/valentina_media/085/222/384_360x240.jpg"),
		},
		&fundaObject{
			address:  "Tweede Atjehstraat 26 II 1094 LG Amsterdam",
			price:    "€ 325.000 k.k.",
			url:      parseURL("https://www.funda.nl/koop/amsterdam/appartement-49389196-tweede-atjehstraat-26-ii/"),
			imageURL: parseURL("https://cloud.funda.nl/valentina_media/085/236/192_360x240.jpg"),
		},
	})

	require.NotNil(t, nextURL)
	require.Equal(t, *nextURL, parseURL("https://www.funda.nl/koop/amsterdam/p2/"))
}

func parseURL(s string) url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	return *u
}
