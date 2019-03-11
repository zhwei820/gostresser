package grpc

import (
	pb "github.com/zhwei820/gostresser/pb/say"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50051"
)

var GrpcSayClient pb.SayServiceClient
var GrpcConn *grpc.ClientConn

func init() {
	// Set up a connection to the server.
	GrpcConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	GrpcSayClient = pb.NewSayServiceClient(GrpcConn)
}

func CloseConn() {
	GrpcConn.Close()
}
