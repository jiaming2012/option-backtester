package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jiaming2012/option-backtester/src/models"
)

func FetchHistOptionOHLC(baseURL string, r models.ThetaDataHistOptionOHLCRequest) (*models.ThetaDataHistOptionOHLCResponse, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	url := fmt.Sprintf("%s/v2/hist/option/ohlc", baseURL)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("FetchHistOptionOHLC: failed to create request: %w", err)
	}

	q := req.URL.Query()
	q.Add("root", string(r.Root))
	q.Add("right", string(r.Right))
	q.Add("exp", r.Expiration.Format("20060102"))
	q.Add("start_date", r.StartDate.Format("20060102"))
	q.Add("end_date", r.EndDate.Format("20060102"))
	q.Add("ivl", fmt.Sprintf("%d", (int(r.Interval/time.Minute)*60000)))
	q.Add("strike", fmt.Sprintf("%d", int(r.Strike*1000.0)))

	req.URL.RawQuery = q.Encode()
	req.Header.Add("Accept", "application/json")

	log.Printf("FetchHistOptionOHLC: fetching option ohlc from %v", req.URL.String())

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("FetchHistOptionOHLC: failed to fetch option ohlc: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("FetchHistOptionOHLC: failed to fetch option ohlc, http code %v", res.Status)
	}

	var dto models.ThetaDataHistOptionOHLCResponse
	if err := json.NewDecoder(res.Body).Decode(&dto); err != nil {
		return nil, fmt.Errorf("FetchHistOptionOHLC: failed to decode json: %w", err)
	}

	return &dto, nil
}
