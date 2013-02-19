package rss

import "testing"

func TestRSSXML(t *testing.T) {
	rss, err := ReadFile("tests/appcast.xml");

	if err != nil {
		t.Errorf("RSS read fail.")
	}

	channel := rss.Channel

	if len(channel.Item) == 0 {
		t.Errorf("Item length is zero")
	}

	for _ , item := range channel.Item {
		if len(item.Title) == 0 {
			t.Errorf("Item Title is empty")
		}
	}

	err = WriteFile("tests/appcast-out.xml",rss)
	if err != nil {
		t.Error(err)
	}

}


