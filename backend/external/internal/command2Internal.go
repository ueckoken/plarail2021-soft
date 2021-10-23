package internal

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	pb "ueckoken/plarail2021-soft-external/spec"
)

const (
	address    = "127.0.0.1:12345"
	timeoutSec = 1
)

// main this is sample func. https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_client/main.go
func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewControlClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), timeoutSec*time.Second)
	defer cancel()
	r, err := c.Command2Internal(ctx, &pb.RequestSync{
		Name:  "sasazuka",
		State: pb.RequestSync_ON,
	})
	if err != nil {
		log.Fatalf("something err: %v", err)
	}
	log.Printf("Response: %s", r.GetResponse())
}
