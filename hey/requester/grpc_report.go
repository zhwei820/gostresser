package requester

import (
	"github.com/zhwei820/gostresser/grpc"
	pb "github.com/zhwei820/gostresser/pb/say"
	"github.com/zhwei820/gostresser/stat"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"time"
)

func sum(ErrorDist map[string]int32) int32 {
	var res int32 = 0
	for _, item := range ErrorDist {
		res += item
	}
	return res
}

func SayHello(report stat.Report, id string, req *http.Request) {

	mrt := 0.0
	avgcl := 0.0
	crps := 0
	rps := 0
	if int(len(report.ResLats)/2) > 0 {
		mrt = report.ResLats[int(len(report.ResLats)/2)]
		avgcl = float64(report.SizeTotal) / float64(len(report.ResLats))

		interval := int((stat.Now() - report.Start) / time.Second)
		if report.Rps > 0 || interval == 0 {
			crps = 0
		} else {
			crps = len(report.ResLats) / interval
		}
	}
	rps = crps
	if report.Rps > 0 {

		rps = int(report.Rps)
	} else {
		report.AvgRes /= float64(len(report.ResLats))
	}
	res := stat.SingleStat{
		AvgContentLength:   avgcl,
		AvgResponseTime:    report.AvgRes,
		MaxResponseTime:    report.ResMax,
		MinResponseTime:    report.ResMin,
		MedianResponseTime: mrt,
		CurrentRps:         int32(crps),
		Rps:                int32(rps),
		Method:             req.Method,
		Name:               req.URL.String(),
		NumRequests:        int32(len(report.ResLats)),
		NumFailures:        sum(report.ErrorDist),
	}

	_, err := grpc.GrpcSayClient.SayHello(context.Background(), &pb.SayInput{
		Id:                 id,
		AvgContentLength:   res.AvgContentLength,
		AvgResponseTime:    res.AvgResponseTime,
		MaxResponseTime:    res.MaxResponseTime,
		MinResponseTime:    res.MinResponseTime,
		MedianResponseTime: res.MedianResponseTime,
		CurrentRps:         res.CurrentRps,
		Rps:                res.Rps,
		Method:             res.Method,
		Name:               res.Name,
		NumRequests:        res.NumRequests,
		NumFailures:        res.NumFailures,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}
