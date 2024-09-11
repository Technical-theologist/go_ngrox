package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type Product struct {
	Id                 int       `json:"id"`
	ProductName        string    `json:"productname"`
	ProductDescription string    `json:"productdescription"`
	Price              float32   `json:"price"`
	Quantity           int       `json:"quantity"`
	Category           string    `json:"category"`
	CreatedDate        time.Time `json:"createddate"`
	// "2006-01-02 15:04:05"
}

func (t *Product) UnmarshalJSON(data []byte) error {
	type Alias Product
	aux := &struct {
		CreatedDate string `json:"createddate"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Attempt to parse the createddate with multiple formats
	var createddate time.Time
	var err error
	formats := []string{
		"2006-01-02",
		"2006-01-02 15:04:05",
	}
	for _, format := range formats {
		createddate, err = time.Parse(format, aux.CreatedDate)
		if err == nil {
			t.CreatedDate = createddate
			return nil
		}
	}

	// Return an error if all parsing attempts fail
	return fmt.Errorf("invalid createddate format: %v, data: %v", err, aux.CreatedDate)
}
