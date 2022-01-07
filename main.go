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
	bp := BaseParser()
	ro := createRequestOptions(RequestOptions{
		ProxyURL: "http://localhost:4780",
		Timeout:  6 * time.Second,
	})

	bp.FetchURL("https://news.ycombinator.com/rss", ro)

	// fp := gofeed.NewParser()
	// feed, _ := fp.ParseURL("https://feeds.feedburner.com/ruanyifeng")
	// fmt.Println(feed.Title)
}

type Parser struct {
	UserAgent string
}
type Feed struct{}
type RequestOptions struct {
	UserAgent string
	ProxyURL  string
	Timeout   time.Duration
}

func createRequestOptions(o RequestOptions) RequestOptions {
	options := RequestOptions{
		UserAgent: "RssRequiem",
	}

	if o.ProxyURL != "" {
		options.ProxyURL = o.ProxyURL
	}
	if o.Timeout != 0 {
		options.Timeout = o.Timeout
	}

	return options
}

func createProxy(proxyURL string) http.Transport {
	URL, err := url.Parse(proxyURL)
	if err != nil {
		fmt.Print(err)
	}

	return http.Transport{
		Proxy: http.ProxyURL(URL),
	}

}

func createClient(options RequestOptions) http.Client {
	transport := createProxy(options.ProxyURL)

	return http.Client{
		Timeout:   options.Timeout,
		Transport: &transport,
	}
}

func createRequest(requestURL string, options RequestOptions) (request *http.Request, err error) {
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(context.Background())
	req.Header.Set("User-Agent", options.UserAgent)

	return req, nil
}

func BaseParser() *Parser {
	p := Parser{}

	return &p
}

func (p *Parser) FetchURL(feedURL string, options RequestOptions) (feed *Feed, err error) {
	req, err := createRequest(feedURL, options)
	if err != nil {
		fmt.Print(err)
	}

	client := createClient(options)

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
