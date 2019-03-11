package grpc

import (
	"github.com/zhwei820/gostresser/stat"
	"log"
	"net"
	"time"

	pb "github.com/zhwei820/gostresser/pb/say"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.SayInput) (*pb.SayOutput, error) {
	println("get SayHello")
	res := &stat.Report{
		Method:              in.Method,
		Name:                in.Name,
		Start:               time.Duration(in.Start),
		AvgTotal:            in.AvgTotal,
		Fastest:             in.Fastest,
		Slowest:             in.Slowest,
		Average:             in.Average,
		Rps:                 in.Rps,
		AvgConn:             in.AvgConn,
		AvgDNS:              in.AvgDNS,
		AvgReq:              in.AvgReq,
		AvgRes:              in.AvgRes,
		AvgDelay:            in.AvgDelay,
		ConnMax:             in.ConnMax,
		ConnMin:             in.ConnMin,
		DnsMax:              in.DnsMax,
		DnsMin:              in.DnsMin,
		ReqMax:              in.ReqMax,
		ReqMin:              in.ReqMin,
		ResMax:              in.ResMax,
		ResMin:              in.ResMin,
		DelayMax:            in.DelayMax,
		DelayMin:            in.DelayMin,
		Lats:                in.Lats,
		ConnLats:            in.ConnLats,
		DnsLats:             in.DnsLats,
		ReqLats:             in.ReqLats,
		ResLats:             in.ResLats,
		DelayLats:           in.DelayLats,
		Offsets:             in.Offsets,
		StatusCodes:         in.StatusCodes,
		Total:               time.Duration(in.Total),
		ErrorDist:           in.ErrorDist,
		StatusCodeDist:      in.StatusCodeDist,
		SizeTotal:           in.SizeTotal,
		SizeReq:             in.SizeReq,
		NumRes:              in.NumRes,
		LatencyDistribution: in.LatencyDistribution,
	}
	stat.Reporters.Store(in.Id, res)
	return &pb.SayOutput{Msg: "Hello " + in.Method}, nil
}

func GrpcServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSayServiceServer(s, &server{})
	s.Serve(lis)
}
