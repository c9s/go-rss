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
	Channel			Channel `xml:"channel"`
	Version			string `xml:"version,attr"`
	// <rss version="2.0" xmlns:sparkle="http://www.andymatuschak.org/xml-namespaces/sparkle"  xmlns:dc="http://purl.org/dc/elements/1.1/">
}

type bytes []byte

func (b bytes) Concat(old2 []byte) []byte {
   newslice := make([]byte, len(b) + len(old2))
   copy(newslice, b)
   copy(newslice[len(b):], old2)
   return newslice
}

func ParseContent(text []byte) (*rss, error) {
	rss := rss{}
	err := xml.Unmarshal(text, &rss)
	if err != nil {
		return nil, err
	}
	return &rss, nil
}

func GetXmlHeader() ([]byte) {
	return []byte("<?xml version=\"1.0\" encoding=\"UTF-8\" ?>\n")
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

func Write(rss * rss) ([]byte,error) {
	rss.Version = "2.0"
	text, err := xml.Marshal(&rss)
	if err != nil {
		return nil,err
	}
	return bytes( GetXmlHeader() ).Concat(text),nil
}

func WriteIndent(rss *rss) ([]byte,error) {
	rss.Version = "2.0"
	text, err := xml.MarshalIndent(&rss,"","  ")
	if err != nil {
		return nil,err
	}
	return bytes( GetXmlHeader() ).Concat(text),nil
}


func WriteFile(file string, rss * rss ) (error) {
	text, err := Write(rss)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(file, text, 0666)
	if err != nil {
		return err
	}
	return nil
}



