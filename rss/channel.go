package rss

type Channel struct {
	Title         string `xml:"title"`
	Link          string `xml:"link"`
	Description   string `xml:"description"`
	Language      string `xml:"language"`
	LastBuildDate Date   `xml:"lastBuildDate"`
	Items         []Item `xml:"item"`
}

func (c * Channel) AddItem(item * Item) {
	c.Items = append(c.Items , *item)
}


