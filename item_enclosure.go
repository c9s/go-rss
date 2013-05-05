package rss

type ItemEnclosure struct {
	URL  string `xml:"url,attr"`
	Type string `xml:"type,attr"`
	Length int64 `xml:"length,attr"`
}

