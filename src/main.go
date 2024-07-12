package main

import (
	"fmt"
	"time"

	"github.com/jiaming2012/option-backtester/src/models"
	"github.com/jiaming2012/option-backtester/src/services"
)

func main() {
	req := models.ThetaDataHistOptionOHLCRequest{
		Root:       "AAPL",
		Right:      models.OptionTypeCall,
		Expiration: time.Date(2023, time.November, 3, 0, 0, 0, 0, time.UTC),
		Strike:     170.0,
		StartDate:  time.Date(2023, time.November, 3, 0, 0, 0, 0, time.UTC),
		EndDate:    time.Date(2023, time.November, 3, 0, 0, 0, 0, time.UTC),
		Interval:   15 * time.Minute,
	}

	resp, err := services.FetchHistOptionOHLC(req)
	if err != nil {
		panic(fmt.Errorf("failed to fetch option ohlc: %w", err))
	}

	fmt.Println(resp)
}
