// Package raw provides functionality for working with raw stock data.

package raw

// Stock represents a stock with its ticker symbol, gap percentage, and opening price.
// This might come from a CSV file.
type Stock struct {
	Ticker       string  // Ticker symbol of the stock
	Gap          float64 // Gap is a percentage of the previous day's closing price
	OpeningPrice float64 // Today's opening price of the stock
}

// Loader is an interface for loading stock data.
type Loader interface {
	Load() ([]Stock, error)
}

// Filterer is an interface for filtering raw stock data.
type Filterer interface {
	Filter([]Stock) []Stock
}
