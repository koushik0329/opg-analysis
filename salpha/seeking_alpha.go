package salpha

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/codebuilds-dev/opg-analysis/stock/news"
)

const (
	urlPath         = "/news/v2/list-by-symbol"
	hostHeader      = "x-rapidapi-host"
	hostHeaderValue = "seeking-alpha.p.rapidapi.com"
	apiKeyHeader    = "x-rapidapi-key"
	pageSize        = 5
	pageNumber      = 1
)

type client struct {
	baseURL string
	apiKey  string
}

func (c *client) Fetch(ticker string) ([]news.Article, error) {
	// for the url, we could have just used:
	// "https://seeking-alpha.p.rapidapi.com/news/v2/list-by-symbol?size=5&number=1&id=" + ticker
	// but this approach is a better practice

	url, err := c.buildURL(ticker)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(hostHeader, hostHeaderValue)
	req.Header.Add(apiKeyHeader, c.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	res := &SeekingAlphaResponse{}
	err = json.Unmarshal(bodyBytes, res)

	if err != nil {
		return nil, err
	}

	var articles []news.Article

	for _, item := range res.Data {
		art := news.Article{
			Ticker:    ticker,
			PublishOn: item.Attributes.PublishOn,
			Headline:  item.Attributes.Title,
		}
		articles = append(articles, art)
	}

	return articles, nil
}

func (c *client) buildURL(ticker string) (string, error) {
	// Parse the base URL
	parsedURL, err := url.Parse(c.baseURL)
	if err != nil {
		return "", err
	}

	// Add the path to the base URL
	parsedURL.Path += urlPath

	// Set query parameters
	params := url.Values{}
	params.Add("size", fmt.Sprint(pageSize))
	params.Add("number", fmt.Sprint(pageNumber))
	params.Add("id", ticker)
	parsedURL.RawQuery = params.Encode()

	return parsedURL.String(), nil
}

func NewClient(baseURL, apiKey string) news.Fetcher {
	return &client{baseURL: baseURL, apiKey: apiKey}
}
