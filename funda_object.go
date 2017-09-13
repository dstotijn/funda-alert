package main

import "net/url"

type fundaObject struct {
	address  string
	price    string
	url      url.URL
	imageURL url.URL
}
