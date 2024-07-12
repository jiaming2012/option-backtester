package models

import "fmt"

type ThetaDataHistOptionOHLCResponse struct {
	Header   ThetaDataResponseHeader `json:"header"`
	Response [][]interface{}         `json:"response"`
}

func (r *ThetaDataHistOptionOHLCResponse) ToHistOptionOhlcDTO() ([]*HistOptionOhlcDTO, error) {
	out := make([]*HistOptionOhlcDTO, 0)

	for _, row := range r.Response {
		msOfDay, ok := row[0].(float64)
		if !ok {
			return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: unable to convert ms_of_day to float64")
		}
		msOfDay2, ok := row[1].(float64)
		if !ok {
			return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: unable to convert ms_of_day2 to float64")
		}
		open, ok := row[2].(float64)
		if !ok {
			return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: unable to convert open to float64")
		}
		high, ok := row[3].(float64)
		if !ok {
			return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: unable to convert high to float64")
		}
		low, ok := row[4].(float64)
		if !ok {
			return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: unable to convert low to float64")
		}
		close, ok := row[5].(float64)
		if !ok {
			return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: unable to convert close to float64")
		}
		volume, ok := row[6].(float64)
		if !ok {
			return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: unable to convert volume to float64")
		}
		date, ok := row[7].(string)
		if !ok {
			return nil, fmt.Errorf("ThetaDataHistOptionOHLCResponse: unable to convert date to string")
		}
		
		out = append(out, &HistOptionOhlcDTO{
			MsOfDay:  int(msOfDay),
			MsOfDay2: int(msOfDay2),
			Open:     open,
			High:     high,
			Low:      low,
			Close:    close,
			Volume:   int(volume),
			Date:     date,
		})
	}

	return out, nil
}
