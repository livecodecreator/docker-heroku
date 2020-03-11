package mackerel

// MetricDataPoint is
type MetricDataPoint struct {
	Name  string  `json:"name"`
	Time  int64   `json:"time"`
	Value float64 `json:"value"`
}

// MetricDataPoints is
type MetricDataPoints []MetricDataPoint
