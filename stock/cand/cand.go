package cand

type Data struct {
	Ticker       string
	Gap          float64
	OpeningPrice float64
}

type Loader interface {
	Load() ([]Data, error)
}
