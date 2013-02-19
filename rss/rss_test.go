package rss

import "testing"

func TestRSSXML(t *testing.T) {
	rss, err := ReadFile("tests/appcast.xml");

	if err != nil {
		t.Errorf("RSS read fail.")
	}

	channel := rss.Channel

	if len(channel.Items) == 0 {
		t.Errorf("Item length is zero")
	}

	for _ , item := range channel.Items {
		if len(item.Title) == 0 {
			t.Errorf("Item Title is empty")
		}
	}

	text, err := WriteIndent(rss)
	if err != nil {
		t.Error(err)
	}
	if len(text) == 0 {
		t.Error("empty output")
	}
	// fmt.Println("%s",string(text))
}


