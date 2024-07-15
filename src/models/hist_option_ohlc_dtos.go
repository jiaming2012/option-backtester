package models

import (
	"fmt"
	"time"
)

type HistOptionOhlcDTOs []*HistOptionOhlcDTO

func (dtos HistOptionOhlcDTOs) ConvertToHistOptionOhlc(loc *time.Location) ([]*HistOptionOhlc, error) {
	result := make([]*HistOptionOhlc, len(dtos))
	for i, dto := range dtos {
		candle, err := dto.ToHistOptionOhlc(loc)
		if err != nil {
			return nil, fmt.Errorf("HistOptionOhlcDTOs.ConvertToHistOptionOhlc: %w", err)
		}

		result[i] = candle
	}

	return result, nil
}
