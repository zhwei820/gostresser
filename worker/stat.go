package worker

type SingleStat struct {
	Method             string  `json:"method" bson:"method"`
	Name               string  `json:"name" bson:"name"`
	NumRequests        int     `json:"num_requests" bson:"num_requests"`
	NumFailures        int     `json:"num_failures" bson:"num_failures"`
	MedianResponseTime float64 `json:"median_response_time" bson:"median_response_time"`
	AvgResponseTime    float64 `json:"avg_response_time" bson:"avg_response_time"`
	MinResponseTime    float64 `json:"min_response_time" bson:"min_response_time"`
	MaxResponseTime    float64 `json:"max_response_time" bson:"max_response_time"`
	AvgContentLength   float64 `json:"avg_content_length" bson:"avg_content_length"`
	CurrentRps         float64 `json:"current_rps" bson:"current_rps"`
}

type StatRes struct {
	CurrentResponseTimePercentile50 int          `json:"current_response_time_percentile_50" bson:"current_response_time_percentile_50"`
	CurrentResponseTimePercentile95 int          `json:"current_response_time_percentile_95" bson:"current_response_time_percentile_95"`
	Errors                          []string     `json:"errors" bson:"errors"`
	FailRatio                       float64      `json:"fail_ratio" bson:"fail_ratio"`
	TotalRps                        int          `json:"total_rps" bson:"total_rps"`
	UserCount                       int          `json:"user_count" bson:"user_count"`
	Stats                           []SingleStat `json:"stats" bson:"stats"`
}

func mean(lats []float64) float64 {
	res := 0.0
	for _, item := range lats {
		res += item
	}
	return res / float64(len(lats))
}
func sum(ErrorDist map[string]int) int {
	res := 0
	for _, item := range ErrorDist {
		res += item
	}
	return res
}

func StatReqs() *StatRes {
	UserCount := 0
	for _, item := range Workers {
		UserCount += item.C
	}
	res := StatRes{
		UserCount: UserCount,
	}
	for _, item := range Workers {
		report := item.Report.Snapshot()
		res.Stats = append(res.Stats, SingleStat{
			AvgContentLength:   float64(report.SizeTotal) / float64(len(report.ResLats)),
			AvgResponseTime:    report.AvgRes,
			MaxResponseTime:    report.ResMax,
			MinResponseTime:    report.ResMin,
			MedianResponseTime: report.ResLats[len(report.ResLats)/2],
			CurrentRps:         report.Rps,
			Method:             item.Request.Method,
			Name:               item.Request.URL.Path,
			NumRequests:        len(report.ResLats),
			NumFailures:        sum(report.ErrorDist),
		})
	}
	return &res
}
