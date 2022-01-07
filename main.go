package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	p := CreateParser()
	p.FetchURL("https://news.ycombinator.com/rss")

	// fp := gofeed.NewParser()
	// feed, _ := fp.ParseURL("https://feeds.feedburner.com/ruanyifeng")
	// fmt.Println(feed.Title)
}

type Parser struct {
	UserAgent string
}
type Feed struct{}

func CreateParser() *Parser {
	p := Parser{
		UserAgent: "RssRequiem",
	}

	return &p
}

func (p *Parser) FetchURL(feedURL string) (feed *Feed, err error) {
	req, err := http.NewRequest("GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(context.Background())
	req.Header.Set("User-Agent", p.UserAgent)

	proxyURL, err := url.Parse("http://localhost:4780")
	transport := http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	client := http.Client{
		Timeout:   6 * time.Second,
		Transport: &transport,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(string(body))

	// TODO placeholder for the real parser
	f := Feed{}
	return &f, nil
}
