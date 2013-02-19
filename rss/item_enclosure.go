package rss

type ItemEnclosure struct {
	URL  string `xml:"url,attr"`
	Type string `xml:"type,attr"`
	Length int  `xml:"length,attr"`
}

