package model

import (
	"encoding/xml"
	"time"
)

// A Feed represents the singular of feed.
type Feed struct {
	XMLName  xml.Name    `xml:"http://www.w3.org/2005/Atom feed"`
	Title    string      `xml:"title"`
	Subtitle string      `xml:"subtitle"`
	Link     FeedLink    `xml:"link"`
	Updated  time.Time   `xml:"updated"`
	Author   FeedAuthor  `xml:"author"`
	ID       string      `xml:"id"`
	Entries  []FeedEntry `xml:"entry"`
}

// An FeedEntry represents the singular of entry.
type FeedEntry struct {
	Title     string      `xml:"title"`
	Link      FeedLink    `xml:"link"`
	ID        string      `xml:"id"`
	Updated   time.Time   `xml:"updated"`
	Published time.Time   `xml:"published"`
	Author    FeedAuthor  `xml:"author"`
	Content   FeedContent `xml:"content"`
}

// An FeedAuthor represents the singluar of author.
type FeedAuthor struct {
	Name string `xml:"name"`
}

// A FeedLink represents the singular of link.
type FeedLink struct {
	Href string `xml:"href,attr"`
}

// A FeedContent represents the singular of content.
type FeedContent struct {
	Type  string `xml:"type,attr"`
	CDATA string `xml:",cdata"`
}
