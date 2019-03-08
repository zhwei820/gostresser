package main

import (
	"log"
	"net"

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
	return &pb.SayOutput{Title: "Hello " + in.Query}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSayServiceServer(s, &server{})
	s.Serve(lis)
}
