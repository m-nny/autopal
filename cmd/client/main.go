package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "minmax.uk/autopal/proto"
)

var (
	addr     = flag.String("addr", "localhost:50000", "address to connect to")
	username = flag.String("name", "kenobi", "name to report")
	timeout  = flag.Duration("timeout", time.Second, "rpc timeout")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMainServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()
	r, err := c.GetUserInfo(ctx, &pb.GetUserInfoRequest{Username: *username})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("response: {%+v}", r)
}
