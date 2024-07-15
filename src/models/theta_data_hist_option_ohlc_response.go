package models

import (
	"fmt"
)

type ThetaDataHistOptionOHLCResponse struct {
	Header   ThetaDataResponseHeader `json:"header"`
	Response [][]interface{}         `json:"response"`
}

func (r *ThetaDataHistOptionOHLCResponse) GetHeaderIndex(headerName string) (int, error) {
	for i, v := range r.Header.Format {
		if v == headerName {
			return i, nil
		}
	}

	return -1, fmt.Errorf("ThetaDataHistOptionOHLCResponse: unable to find header %v", headerName)
}

func (r *ThetaDataHistOptionOHLCResponse) ToHistOptionOhlcDTO() ([]*HistOptionOhlcDTO, error) {
	out := make([]*HistOptionOhlcDTO, 0)

	msOfDayIndex, err := r.GetHeaderIndex("ms_of_day")
	if err != nil {
		return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: %w", err)
	}

	openIndex, err := r.GetHeaderIndex("open")
	if err != nil {
		return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: %w", err)
	}

	highIndex, err := r.GetHeaderIndex("high")
	if err != nil {
		return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: %w", err)
	}

	lowIndex, err := r.GetHeaderIndex("low")
	if err != nil {
		return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: %w", err)
	}

	closeIndex, err := r.GetHeaderIndex("close")
	if err != nil {
		return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: %w", err)
	}

	volumeIndex, err := r.GetHeaderIndex("volume")
	if err != nil {
		return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: %w", err)
	}

	dateIndex, err := r.GetHeaderIndex("date")
	if err != nil {
		return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: %w", err)
	}

	for _, row := range r.Response {
		msOfDay, ok := row[msOfDayIndex].(float64)
		if !ok {
			return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: msOfDay, unable to convert ms_of_day to float64, %v", row[0])
		}
		open, ok := row[openIndex].(float64)
		if !ok {
			return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: open, unable to convert open to float64, %v", row[2])
		}
		high, ok := row[highIndex].(float64)
		if !ok {
			return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: high, unable to convert high to float64, %v", row[3])
		}
		low, ok := row[lowIndex].(float64)
		if !ok {
			return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: low, unable to convert low to float64, %v", row[4])
		}
		close, ok := row[closeIndex].(float64)
		if !ok {
			return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: close, unable to convert close to float64, %v", row[5])
		}
		volume, ok := row[volumeIndex].(float64)
		if !ok {
			return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: volume, unable to convert volume to float64, %v", row[6])
		}
		date, ok := row[dateIndex].(float64)
		if !ok {
			return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: date, unable to convert date to string, %v", row[7])
		}

		out = append(out, &HistOptionOhlcDTO{
			MsOfDay: int(msOfDay),
			Open:    open,
			High:    high,
			Low:     low,
			Close:   close,
			Volume:  int(volume),
			Date:    int(date),
		})
	}

	return out, nil
}
