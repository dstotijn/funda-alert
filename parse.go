package main

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func fundaObjects(r io.Reader) (objects []*fundaObject, nextURL *url.URL, err error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return
	}

	doc.Find("link").Each(func(i int, s *goquery.Selection) {
		rel, relExists := s.Attr("rel")
		href, hrefExists := s.Attr("href")
		if relExists && rel == "next" && hrefExists {
			nextURL, err = url.Parse(href)
			if err != nil {
				err = fmt.Errorf("could not parse `next` URL: %s", err)
				return
			}
		}
	})

	doc.Find(".search-result").Each(func(i int, s *goquery.Selection) {
		houseURL, err := url.Parse("https://www.funda.nl" + s.Find(".search-result-header a").AttrOr("href", ""))
		if err != nil {
			log.Printf("Error parsing house URL: %s", err)
			return
		}
		imageURL, err := url.Parse(s.Find(".search-result-image img").AttrOr("src", ""))
		if err != nil {
			log.Printf("Error parsing house URL: %s", err)
			return
		}
		h := &fundaObject{
			address:  cleanText(s.Find(".search-result-title").Text()),
			price:    cleanText(s.Find(".search-result-price").Text()),
			url:      *houseURL,
			imageURL: *imageURL,
		}
		objects = append(objects, h)
	})

	return
}

func cleanText(s string) string {
	// Strip extraneous whitespace chars (including newlines).
	return strings.Join(strings.Fields(s), " ")
}
