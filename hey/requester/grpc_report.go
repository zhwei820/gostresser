package requester

import (
	"fmt"
	pb "github.com/zhwei820/gostresser/pb/say"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func GrpcSayHello() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSayServiceClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.SayInput{Query: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Printf("\nGreeting: %s\n", r.Title)
}
