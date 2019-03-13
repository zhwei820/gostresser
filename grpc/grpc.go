package grpc

import (
	pb "github.com/zhwei820/gostresser/pb/say"
	"github.com/zhwei820/gostresser/stat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.SayInput) (*pb.SayOutput, error) {
	res := &stat.SingleStat{
		Id:                 in.Id,
		AvgContentLength:   in.AvgContentLength,
		AvgResponseTime:    in.AvgResponseTime,
		MaxResponseTime:    in.MaxResponseTime,
		MinResponseTime:    in.MinResponseTime,
		MedianResponseTime: in.MedianResponseTime,
		CurrentRps:         in.CurrentRps,
		Rps:                in.Rps,
		Method:             in.Method,
		Name:               in.Name,
		NumRequests:        in.NumRequests,
		NumFailures:        in.NumFailures,
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
