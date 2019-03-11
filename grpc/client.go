package grpc

import (
	"fmt"
	"github.com/zhwei820/gostresser/hey/requester"
	pb "github.com/zhwei820/gostresser/pb/say"
	"golang.org/x/net/context"
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

func SayHello(report requester.Report) {

	r, err := GrpcSayClient.SayHello(context.Background(), &pb.SayInput{Name: "name"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("\nGreeting: %s\n", r.Msg)
}

func CloseConn() {
	GrpcConn.Close()
}
