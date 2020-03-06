package data

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// TODO: add comment
func DataSetRequest() (*resty.Response, error) {
	client := resty.New()

	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		Get("https://data.austintexas.gov/api/views/c2a6-efzi/rows.json")
	if err != nil {
		return nil, fmt.Errorf("Error while retrieving dataset: %v", err)
	}

	return res, nil
}
