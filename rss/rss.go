/*
Simple RSS parser, tested with Wordpress feeds.
*/
package rss

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

var rss struct {
	Channel Channel `xml:"channel"`
}

func ParseContent(text []byte) (*Channel, error) {
	err := xml.Unmarshal(text, &rss)
	if err != nil {
		return nil, err
	}
	return &rss.Channel, nil
}

func ReadFile(file string) (*Channel, error) {
	text, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return ParseContent(text)
}

func ReadUrl(url string) (*Channel, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	text, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return ParseContent(text)
}


