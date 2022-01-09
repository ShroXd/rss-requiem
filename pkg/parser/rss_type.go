package parser

import "encoding/xml"

type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title          string `xml:"title"`
	Link           string `xml:"link"`
	Description    string `xml:"description"`
	Language       string `xml:"language,omitempty"`
	Copyright      string `xml:"copyright"`
	ManagingEditor string `xml:"managingEditor"`
	PubDate        string `xml:"pubDate"`
	LastBuildDate  string `xml:"lastBuildDate"`
	Generator      string `xml:"generator"`
	Docs           string `xml:"docs"`
	Cloud          string `xml:"cloud"`
	Ttl            string `xml:"ttl"`
	Image          string `xml:"image"`
	TextInput      string `xml:"textInput"`
	SkipHours      string `xml:"skipHours"`
	SkipDays       string `xml:"skipDays"`
	Item           []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Author      string `xml:"author"`
	Category    string `xml:"category"`
	Comments    string `xml:"comments"`
	Enclosure   string `xml:"enclosure"`
	Guid        string `xml:"guid"`
	PubDate     string `xml:"pubDate"`
	Source      string `xml:"source"`
}
