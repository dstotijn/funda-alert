package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type fundaObject struct {
	id          int
	address     string
	price       string
	url         url.URL
	imageURL    url.URL
	surfaceArea string
	rooms       string
}

type fundaObjects []*fundaObject

type foto struct {
	Link string `json:"Link"`
}

type line struct {
	Text string `json:"Text"`
}

type info struct {
	Line []line `json:"Line"`
}

type fundaSearchResultItem struct {
	ItemType int `json:"ItemType`
	Fotos    []foto `json:"Fotos"`
	Info     []info `json:"Info"`
	GlobalID int    `json:"GlobalId"`
}

type fundaSearchResult []fundaSearchResultItem

func searchFunda(fundaToken, fundaSearchOptions string, page, pageSize int) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("accepted_cookie_policy", "10")
	req.Header.Set("api_key", fundaToken)
	req.Header.Set("User-Agent", "Funda/2.17.0 (com.funda.two; build:80; Android 25) okhttp/3.5.0")
	req.Header.Set("Cookie", "X-Stored-Data=null; expires=Fri, 31 Dec 9999 23:59:59 GMT; path=/; samesite=lax; httponly")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "nl-NL")

	u, err := fundaSearchURL(fundaSearchOptions, page, pageSize)
	if err != nil {
		return nil, err
	}
	req.URL = u

	log.Printf("Searching Funda for objects: %v", u.String())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected HTTP response code `%d` received", resp.StatusCode)
	}

	return resp.Body, nil
}

func fundaObjectsFromSearchResult(r io.Reader) (objects fundaObjects, err error) {
	var result fundaSearchResult
	if err = json.NewDecoder(r).Decode(&result); err != nil {
		return
	}

	for _, o := range result {
		// Skip highlighted funda objects (ads).
		if o.ItemType != 1 {
			continue
		}
		if len(o.Fotos) < 1 {
			return nil, errors.New("result does not have photos")
		}
		if len(o.Info) < 4 {
			return nil, errors.New("result does not have enough info values")
		}

		for _, info := range o.Info {
			if len(info.Line) < 1 {
				return nil, errors.New("result does not have enough info lines")
			}
		}

		var houseURL, imageURL *url.URL
		object := &fundaObject{
			id:      o.GlobalID,
			address: o.Info[0].Line[0].Text,
			price:   o.Info[3].Line[0].Text,
		}

		surfaceAreaRooms := strings.Split(o.Info[2].Line[0].Text, " • ")
		if len(surfaceAreaRooms) != 2 {
			return nil, fmt.Errorf("unexpected surface area & price line: %v", surfaceAreaRooms)
		}

		object.surfaceArea = surfaceAreaRooms[0]
		object.rooms = surfaceAreaRooms[1]

		houseURL, err = url.Parse(fmt.Sprintf("https://www.funda.nl/%v", o.GlobalID))
		if err != nil {
			log.Printf("Error parsing house URL: %s", err)
			return
		}
		object.url = *houseURL

		imageURL, err = url.Parse(o.Fotos[0].Link)
		if err != nil {
			log.Printf("Error parsing image URL: %s", err)
			return
		}
		object.imageURL = *imageURL

		objects = append(objects, object)
	}

	return
}

func fundaSearchURL(searchOptions string, page, pageSize int) (*url.URL, error) {
	u, err := url.Parse("https://mobile.funda.io/api/v1/Aanbod/koop" + searchOptions)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Set("page", strconv.Itoa(page))
	q.Set("pageSize", strconv.Itoa(pageSize))

	u.RawQuery = q.Encode()

	return u, nil
}
