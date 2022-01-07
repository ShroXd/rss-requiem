package request

import (
	"atriiy/rss-requiem/pkg/parser"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type RequestOptions struct {
	UserAgent string
	ProxyURL  string
	Timeout   time.Duration
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

func CreateRequestOptions(o RequestOptions) RequestOptions {
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

func createRequest(requestURL string, options RequestOptions) (request *http.Request, err error) {
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(context.Background())
	req.Header.Set("User-Agent", options.UserAgent)

	return req, nil
}

func createClient(options RequestOptions) http.Client {
	transport := createProxy(options.ProxyURL)

	return http.Client{
		Timeout:   options.Timeout,
		Transport: &transport,
	}
}

func FetchURL(feedURL string, options RequestOptions) (feed *parser.Feed, err error) {
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
	f := parser.Feed{}
	return &f, nil
}
