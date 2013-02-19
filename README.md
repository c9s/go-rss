Go RSS Parser
=============

Install
--------

    go get github.com/c9s/go-rss/rss


Usage
-----

    import "github.com/c9s/go-rss/rss"
    r, err := rss.ReadUrl("http://....path/to/feed.xml")

    err = rss.WriteFile("feedout.xml",r)


License
-------

Public Domain
