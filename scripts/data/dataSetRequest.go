package data

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// MakeRequest retrieves data from Data.gov dataset
// using resty client
func MakeRequest() (*resty.Response, error) {
	client := resty.New()

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		Get("https://data.austintexas.gov/api/views/c2a6-efzi/rows.json")
	if err != nil {
		return nil, fmt.Errorf("Error while retrieving dataset: %v", err)
	}

	return res, nil
}
