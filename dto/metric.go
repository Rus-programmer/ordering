package dto

type MetricsResponse struct {
	TotalRequests    int64            `json:"total_requests"`
	RequestsByMethod map[string]int64 `json:"requests_by_method"`
	RequestsByStatus map[int32]int64  `json:"requests_by_status_code"`
	RequestsByPath   map[string]int64 `json:"requests_by_path"`
	ErrorRate        float64          `json:"error_rate"`
}
