package models

import (
	"fmt"
	"strconv"
	"time"
)

type HistOptionOhlcDTO struct {
	MsOfDay int     `json:"ms_of_day"`
	Open    float64 `json:"open"`
	High    float64 `json:"high"`
	Low     float64 `json:"low"`
	Close   float64 `json:"close"`
	Volume  int     `json:"volume"`
	Date    int     `json:"date"`
}

// ConvertDateAndMsToTime takes a date in YYYYMMDD format and milliseconds since midnight and returns a time.Time object
func convertDateAndMsToTime(date int, msOfDay int, loc *time.Location) (time.Time, error) {
	// Parse the date
	dateStr := strconv.Itoa(date)
	year, err := strconv.Atoi(dateStr[:4])
	if err != nil {
		return time.Time{}, fmt.Errorf("convertDateAndMsToTime: failed to convert date to year: %w", err)
	}
	month, err := strconv.Atoi(dateStr[4:6])
	if err != nil {
		return time.Time{}, fmt.Errorf("convertDateAndMsToTime: failed to convert date to month: %w", err)
	}
	day, err := strconv.Atoi(dateStr[6:])
	if err != nil {
		return time.Time{}, fmt.Errorf("convertDateAndMsToTime: failed to convert date to day: %w", err)
	}

	// Convert milliseconds to hours, minutes, and seconds
	seconds := msOfDay / 1000
	msRemaining := msOfDay % 1000
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	secs := seconds % 60

	// Combine the date and time components
	return time.Date(year, time.Month(month), day, hours, minutes, secs, msRemaining*1e6, loc), nil
}

func (dto *HistOptionOhlcDTO) ToHistOptionOhlc(loc *time.Location) (*HistOptionOhlc, error) {
	timestamp, err := convertDateAndMsToTime(dto.Date, dto.MsOfDay, loc)
	if err != nil {
		return nil, fmt.Errorf("HistOptionOhlcDTO.ToHistOptionOhlc: failed to convert date and ms to time: %w", err)
	}

	return &HistOptionOhlc{
		Timestamp: timestamp,
		Open:      dto.Open,
		High:      dto.High,
		Low:       dto.Low,
		Close:     dto.Close,
		Volume:    dto.Volume,
	}, nil
}


