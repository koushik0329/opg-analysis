package news

import "time"

type Article struct {
	Ticker    string
	PublishOn time.Time
	Headline  string
}

type Fetcher interface {
	Fetch(ticker string) ([]Article, error)
}
