package rss

import "testing"

func TestDate(t *testing.T) {
	date := Date("Wed, 09 Jan 2006 19:20:11 +0000")
	time, err := date.Parse()
	if err != nil {
		t.Error(err)
	}
	_ = time
}

