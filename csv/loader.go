package csv

import (
	"encoding/csv"
	"os"
	"slices"
	"strconv"

	"github.com/codebuilds-dev/opg-analysis/stock/cand"
)

type columns = []string
type rows = []columns

type loader struct {
	path string
}

func (l *loader) Load() ([]cand.Data, error) {
	rows, err := l.read()
	if err != nil {
		return nil, err
	}

	var data []cand.Data

	for _, row := range rows {
		ticker := row[0]
		gap, err := strconv.ParseFloat(row[1], 64)

		if err != nil {
			continue
		}

		openingPrice, err := strconv.ParseFloat(row[2], 64)

		if err != nil {
			continue
		}

		data = append(data, cand.Data{
			Ticker:       ticker,
			Gap:          gap,
			OpeningPrice: openingPrice,
		})
	}

	return data, nil
}

func (l *loader) read() (rows, error) {
	f, err := os.Open(l.path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	rows, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Equivalent t this records = append(rows[:0], rows[1:]...)
	rows = slices.Delete(rows, 0, 1)

	return rows, nil
}

func NewLoader(path string) cand.Loader {
	return &loader{
		path: path,
	}
}
