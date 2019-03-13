package stat

import (
	"sync"
)

var Reporters sync.Map

type SingleStat struct {
	Id                 string  `json:"id" bson:"id"`
	Method             string  `json:"method" bson:"method"`
	Name               string  `json:"name" bson:"name"`
	NumRequests        int32   `json:"num_requests" bson:"num_requests"`
	NumFailures        int32   `json:"num_failures" bson:"num_failures"`
	MedianResponseTime float64 `json:"median_response_time" bson:"median_response_time"`
	AvgResponseTime    float64 `json:"avg_response_time" bson:"avg_response_time"`
	MinResponseTime    float64 `json:"min_response_time" bson:"min_response_time"`
	MaxResponseTime    float64 `json:"max_response_time" bson:"max_response_time"`
	AvgContentLength   float64 `json:"avg_content_length" bson:"avg_content_length"`
	CurrentRps         int32   `json:"current_rps" bson:"current_rps"`
	Rps                int32   `json:"rps" bson:"rps"`
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
func sum(ErrorDist map[string]int32) int32 {
	var res int32 = 0
	for _, item := range ErrorDist {
		res += item
	}
	return res
}

func StatReqs() *StatRes {

	res := StatRes{}
	Reporters.Range(func(key, value interface{}) bool {
		sstat := value.(*SingleStat)

		res.Stats = append(res.Stats, *sstat)
		return true
	})
	return &res
}
