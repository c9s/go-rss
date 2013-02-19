package rss

type Item struct {
	Title       string        `xml:"title"`
	Link        string        `xml:"link,omitempty"`
	Comments    string        `xml:"comments,omitempty"`
	PubDate     Date          `xml:"pubDate"`
	GUID        string        `xml:"guid,omitempty"`
	Category    []string      `xml:"category,omitempty"`
	Enclosure   ItemEnclosure `xml:"enclosure"`
	Description string        `xml:"description"`
	Content     string        `xml:"content,omitempty"`
}

func (i *Item) AddCategory (c string) {
	i.Category = append(i.Category,c)
}

func (item * Item) SetEnclosure(enclosure * ItemEnclosure) {
	item.Enclosure = *enclosure
}


