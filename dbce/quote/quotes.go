package dbcapiego

import (
	"encoding/json"

	"fmt"
)

// FixedQuote represents the individual fixed quotes returned by the DBCE API.
type FixedQuote struct {
	Id           string      `json:"id,omitempty"`
	Platform     *Platform   `json:"platform,omitempty"`
	HourlyPrices *Quantities `json:"hourlyPrices,omitempty"`
	TotalPrice   float32     `json:"totalPrice,omitempty"`
}

// Fairly simple for the moment, matching the api.
type Platform struct {
	ID int `json:"id,omitempty"`
}

// Same struct for both HourlyPrices and Quantities.
type Quantities struct {
	Compute float32 `json:"compute,omitempty"`
	Storage float32 `json:"storage,omitempty"`
	Memory  float32 `json:"memory,omitempty"`
}

// The request struct to be sent for a quote request
type FixedQuoteRequest struct {
	Interval   *Interval   `json:"interval,omitempty"`
	Quantities *Quantities `json:"quantities,omitempty"`
	OsTypes    []string    `json:"ostypes,omitempty"`
}

type Interval struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

type AcceptQuoteRequest struct {
	Id string `json:"id,omitempty"`
}
