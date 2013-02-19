/*
Simple RSS parser, tested with Wordpress feeds.
*/
package rss

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type rss struct {
	Channel Channel `xml:"channel"`
}

func ParseContent(text []byte) (*rss, error) {
	rss := rss{}
	err := xml.Unmarshal(text, &rss)
	if err != nil {
		return nil, err
	}
	return &rss, nil
}

func ReadFile(file string) (*rss, error) {
	text, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return ParseContent(text)
}

func ReadUrl(url string) (*rss, error) {
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

func WriteFile(file string, rss * rss ) (error) {
	text, err := xml.Marshal(&rss)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(file, text, 0666)
	if err != nil {
		return err
	}
	return nil
}



