package main

import (
	"atriiy/rss-requiem/pkg/request"
	"time"
)

func main() {
	ro := request.CreateRequestOptions(request.RequestOptions{
		ProxyURL: "http://localhost:4780",
		Timeout:  6 * time.Second,
	})

	// request.FetchURL("https://news.ycombinator.com/rss", ro)
	request.FetchURL("https://feeds.feedburner.com/ruanyifeng", ro)

	// fp := gofeed.NewParser()
	// feed, _ := fp.ParseURL("https://feeds.feedburner.com/ruanyifeng")
	// fmt.Println(feed.Title)
}
