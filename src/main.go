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
		StartDate:  time.Date(2023, time.November, 1, 0, 0, 0, 0, time.UTC),
		EndDate:    time.Date(2023, time.November, 2, 0, 0, 0, 0, time.UTC),
		Interval:   1 * time.Minute,
	}

	baseURL := "http://localhost:25510"
	resp, err := services.FetchHistOptionOHLC(baseURL, req)
	if err != nil {
		panic(fmt.Errorf("failed to fetch option ohlc: %w", err))
	}

	candlesDTO, err := resp.ToHistOptionOhlcDTO()
	if err != nil {
		panic(fmt.Errorf("failed to convert response to dto: %w", err))
	}

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(fmt.Errorf("failed to load location: %w", err))
	}

	candles, err := models.HistOptionOhlcDTOs(candlesDTO).ConvertToHistOptionOhlc(loc)
	if err != nil {
		panic(fmt.Errorf("failed to convert dto to candle: %w", err))
	}

	for i, candle := range candles {
		fmt.Printf("%d: %+v\n", i, candle)
	}
}
