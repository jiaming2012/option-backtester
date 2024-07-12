package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jiaming2012/option-backtester/models"
)

func FetchHistOptionOHLC(req models.HistOptionOHLCRequest) (*models.HistOptionOHLCResponse, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	url := "https://api.thetadatapoint.com/api/v1/option/ohlc"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("FetchHistOptionOHLC: failed to create request: %w", err)
	}

	q := req.URL.Query()
	q.Add("root", req.Root)
	q.Add("right", req.Right)
	q.Add("exp", req.Expiration.Format("20060102"))
	q.Add("start_date", req.StartDate.Format("20060102"))
	q.Add("end_date", req.EndDate.Format("20060102"))
	q.Add("ivl", fmt.Sprintf("%d", req.Interval/1000.0))
	q.Add("strike", fmt.Sprintf("%d", req.Strike*1000.0))

	req.URL.RawQuery = q.Encode()
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("FetchHistOptionOHLC: failed to fetch option ohlc: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("FetchHistOptionOHLC: failed to fetch option ohlc, http code %v", res.Status)
	}

	var dto models.HistOptionOHLCResponse
	if err := json.NewDecoder(res.Body).Decode(&dto); err != nil {
		return nil, fmt.Errorf("FetchHistOptionOHLC: failed to decode json: %w", err)
	}

	if len(dto) == 0 {
		return nil, fmt.Errorf("FetchHistOptionOHLC: no data returned")
	}

	return dto, nil
}
