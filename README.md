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

for _, item := r.Channel.Items {
    // item.Title
    // item.Content
}

newItem := rss.Item{}
newItem.Title = "New Title"
newItem.Content = `Content..............`
r.Channel.AddItem(&newItem)

err = rss.WriteFile("feedout.xml",r)
```


License
-------

Public Domain
