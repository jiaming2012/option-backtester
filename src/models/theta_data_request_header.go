package models

type ThetaDataResponseHeader struct {
	Format    []string `json:"format"`
	NextPage  string   `json:"next_page"`
	LatencyMs int      `json:"latency_ms"`
}
