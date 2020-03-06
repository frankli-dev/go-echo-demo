package data

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/ivan-marquez/golang-demo/pkg/models"
)

// TODO: add comment
func retrieveColumns(value []byte, dataType jsonparser.ValueType, offset int, err error) {
	name, _, _, err := jsonparser.Get(value, "name")

	fmt.Println(string(name))
}

// TODO: add comment
func parseColumnValues(data []byte) ([]*models.RenewableResource, error) {
	var rows []*models.RenewableResource

	parse := func(data []byte, dataType jsonparser.ValueType, offset int, err error) {
		d, _, _, err := jsonparser.Get(data)

		values := cleanupColumnValues(d)

		doc := parseRenewableResource(values)
		rows = append(rows, doc)
	}

	_, err := jsonparser.ArrayEach(data, parse, "data")
	if err != nil {
		// TODO: return error
		return nil, fmt.Errorf("Error parsing data:%v", err)
	}

	return rows, nil
}

// cleanupColumnValues replaces characters
// and trims whitespaces from a byte slice
func cleanupColumnValues(data []byte) []string {
	var c []string

	s := string(data)
	s = strings.Replace(s, "[", "", 1)
	s = strings.Replace(s, "]", "", 1)
	s = strings.ReplaceAll(s, "\"", "")
	s = strings.TrimSpace(s)

	// keep only the necessary values
	values := strings.Split(s, ",")[8:]
	for _, v := range values {
		c = append(c, strings.TrimSpace(v))
	}

	return c
}

// TODO: add comment
func parseRenewableResource(values []string) *models.RenewableResource {
	cd, err := time.Parse("2006-01-02T15:04:05", values[0])
	if err != nil {
		log.Fatalf("Error parsing date: %v", err)
	}

	toFloat := func(value string) float64 {
		var f float64
		if value != "null" {
			f, err = strconv.ParseFloat(strings.TrimSpace(value), 32)
			if err != nil {
				log.Fatalf("Error parsing float: %v", err)
			}
		}

		return f
	}

	toInt := func(value string) int64 {
		var i int64

		if value != "null" {
			i, err = strconv.ParseInt(value, 10, 0)
			if err != nil {
				log.Fatalf("Error parsing int: %v", err)
			}
		}

		return i
	}

	return &models.RenewableResource{
		CalendarDate:                  cd,
		TotalRenewableEnergyResources: float32(toFloat(values[1])),
		InstalledSolarCapacity:        float32(toFloat(values[2])),
		TotalRenewableEnergyPurchased: toInt(values[3]),
		GreenChoiceSales:              toInt(values[4]),
		RenewableEnergyToFuelCharge:   toInt(values[5]),
		Wind:                          float32(toFloat(values[6])),
		UtilityScaleSolar:             float32(toFloat(values[7])),
		Biomass:                       float32(toFloat(values[8])),
	}
}

// TODO: add comment
func ParseResponse(data []byte) ([]*models.RenewableResource, error) {
	// TODO: what to do with description?
	// TODO: what to do with columns?
	description, _, _, err := jsonparser.Get(data, "meta", "view", "description")
	if err != nil {
		return nil, fmt.Errorf("Error while parsing description: %v", err)
	}
	fmt.Println(string(description))

	_, err = jsonparser.ArrayEach(data, retrieveColumns, "meta", "view", "columns")
	if err != nil {
		return nil, fmt.Errorf("Error while parsing columns: %v", err)
	}

	c, err := parseColumnValues(data)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing data values: %v", err)
	}

	return c, nil
}
