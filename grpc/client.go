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
	maxMsgSize := 1024 * 1024 * 100
	GrpcConn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize), grpc.MaxCallSendMsgSize(maxMsgSize)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	GrpcSayClient = pb.NewSayServiceClient(GrpcConn)
}

func CloseConn() {
	GrpcConn.Close()
}
