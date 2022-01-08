package parser

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Parser struct{}

func BaseParser() *Parser {
	p := Parser{}

	return &p
}

type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
}

func (p *Parser) ParseFeed(feed []byte) {
	f1 := Feed{}
	fmt.Println(string(feed), '\n')
	err := xml.Unmarshal(feed, &f1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(f1.XMLName.Local)
	fmt.Println(f1.Version)
}
