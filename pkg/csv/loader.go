// Package csv provides functionality for loading and processing CSV files containing stock data.

package csv

import (
	"encoding/csv"
	"log"
	"os"
	"slices"
	"strconv"

	"github.com/codebuilds-dev/opg-analysis/internal/raw"
)

// columns represents a slice of strings, where each string represents a column in the CSV file.
type columns = []string

// rows represents a slice of columns, where each columns represents a row in the CSV file.
type rows = []columns

// loader is responsible for loading stock data from a CSV file.
type loader struct {
	path string
}

// Load reads the CSV file and returns a slice of raw.Stock objects representing the stock data.
// It skips rows with invalid data and returns an error if there was a problem reading the file.
func (l *loader) Load() ([]raw.Stock, error) {
	rows, err := l.read()
	if err != nil {
		return nil, err
	}

	var data []raw.Stock

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

		data = append(data, raw.Stock{
			Ticker:       ticker,
			Gap:          gap,
			OpeningPrice: openingPrice,
		})
	}

	return data, nil
}

// read reads the CSV file and returns a slice of rows representing the data.
// It removes the header row before returning the rows.
// It returns an error if there was a problem reading the file.
func (l *loader) read() (rows, error) {

	f, err := os.Open(l.path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Remove the header row
	// Equivalent to: rows = slices.Delete(rows, 0, 1)
	rows = slices.Delete(rows, 0, 1)

	log.Printf("Loaded %d rows from %s \n", len(rows), l.path)
	return rows, nil
}

// NewLoader creates a new instance of the loader with the specified file path.
func NewLoader(path string) raw.Loader {
	return &loader{
		path: path,
	}
}
