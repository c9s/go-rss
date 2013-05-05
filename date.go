package rss

import "time"

type Date string

func (self Date) Parse() (time.Time, error) {
	t, err := time.Parse(time.RFC1123Z, string(self)) // Wordpress format
	if err != nil {
		t, err = time.Parse(time.RFC822, string(self)) // RSS 2.0 spec
	}
	return t, err
}

func (self Date) Format(format string) (string, error) {
	t, err := self.Parse()
	if err != nil {
		return "", err
	}
	return t.Format(format), nil
}

func (self Date) MustFormat(format string) string {
	s, err := self.Format(format)
	if err != nil {
		return err.Error()
	}
	return s
}

