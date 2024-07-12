package models

import "time"

type ThetaDataHistOptionOHLCRequest struct {
	Expiration time.Time     `json:"exp"`
	Interval   time.Duration `json:"ivl"`
	Right      OptionType    `json:"right"`
	Root       StockSymbol   `json:"root"`
	StartDate  time.Time     `json:"start_date"`
	EndDate    time.Time     `json:"end_date"`
	Strike     float64       `json:"strike"`
}
