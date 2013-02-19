package rss

import "testing"
import "fmt"

func TestRSSXML(t *testing.T) {
	rss, err := ReadFile("tests/appcast.xml");

	if err != nil {
		t.Errorf("RSS read fail.")
	}

	if len(rss.Channel.Items) == 0 {
		t.Errorf("Item length is zero")
	}

	for _ , item := range rss.Channel.Items {
		if len(item.Title) == 0 {
			t.Errorf("Item Title is empty")
		}
	}

	newItem := Item{}
	newItem.Title = "New Title"
	newItem.Content = "New Content"
	rss.Channel.AddItem(&newItem)

	text, err := WriteIndent(rss)
	if err != nil {
		t.Error(err)
	}
	if len(text) == 0 {
		t.Error("empty output")
	}
	fmt.Println("%s",string(text))
	_ = text
}


