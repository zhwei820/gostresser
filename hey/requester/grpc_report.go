package requester

import (
	"fmt"
	"github.com/zhwei820/gostresser/grpc"
	pb "github.com/zhwei820/gostresser/pb/say"
	"github.com/zhwei820/gostresser/stat"
	"golang.org/x/net/context"
	"log"
)

func SayHello(report stat.Report, id string) {

	r, err := grpc.GrpcSayClient.SayHello(context.Background(), &pb.SayInput{Id: id,
		Method:              report.Method,
		Name:                report.Name,
		Start:               int64(report.Start),
		AvgTotal:            report.AvgTotal,
		Fastest:             report.Fastest,
		Slowest:             report.Slowest,
		Average:             report.Average,
		Rps:                 report.Rps,
		AvgConn:             report.AvgConn,
		AvgDNS:              report.AvgDNS,
		AvgReq:              report.AvgReq,
		AvgRes:              report.AvgRes,
		AvgDelay:            report.AvgDelay,
		ConnMax:             report.ConnMax,
		ConnMin:             report.ConnMin,
		DnsMax:              report.DnsMax,
		DnsMin:              report.DnsMin,
		ReqMax:              report.ReqMax,
		ReqMin:              report.ReqMin,
		ResMax:              report.ResMax,
		ResMin:              report.ResMin,
		DelayMax:            report.DelayMax,
		DelayMin:            report.DelayMin,
		Lats:                report.Lats,
		ConnLats:            report.ConnLats,
		DnsLats:             report.DnsLats,
		ReqLats:             report.ReqLats,
		ResLats:             report.ResLats,
		DelayLats:           report.DelayLats,
		Offsets:             report.Offsets,
		StatusCodes:         report.StatusCodes,
		Total:               int32(report.Total),
		ErrorDist:           report.ErrorDist,
		StatusCodeDist:      report.StatusCodeDist,
		SizeTotal:           report.SizeTotal,
		SizeReq:             report.SizeReq,
		NumRes:              report.NumRes,
		LatencyDistribution: report.LatencyDistribution,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("\nGreeting: %s\n", r.Msg)
}
