package models

type HistOptionOhlcDTO struct {
	MsOfDay  int     `json:"ms_of_day"`
	MsOfDay2 int     `json:"ms_of_day2"`
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	Volume   int     `json:"volume"`
	Date     string  `json:"date"`
}
