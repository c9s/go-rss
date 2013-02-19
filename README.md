Go RSS Parser
=============

Install
--------

    go get github.com/c9s/go-rss/rss


Usage
-----


```go
import "github.com/c9s/go-rss/rss"
r, err := rss.ReadUrl("http://....path/to/feed.xml")

// r.Channel.Items

err = rss.WriteFile("feedout.xml",r)
```


License
-------

Public Domain
