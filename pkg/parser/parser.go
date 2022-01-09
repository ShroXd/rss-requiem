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

func (p *Parser) ParseFeed(feed []byte) {
	f1 := Feed{}
	err := xml.Unmarshal(feed, &f1)
	if err != nil {
		log.Fatal(err)
	}

	ci := f1.Channel.Item[0]

	fmt.Println("Item title: ", ci.Title)
	fmt.Println("Item link: ", ci.Link)
	fmt.Println("Item description: ", ci.Description)
}
