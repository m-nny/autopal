package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"minmax.uk/autopal/pkg/life"
	pb "minmax.uk/autopal/proto/life"
)

var (
	addr    = flag.String("addr", "localhost:50001", "address to connect to")
	timeout = flag.Duration("timeout", time.Second, "rpc timeout")

	rows = flag.Int64("rows", 10, "height of the board")
	cols = flag.Int64("cols", 20, "width of the board")
	seed = flag.Int64("seed", 42, "seed for init state")

	iters = flag.Int("iters", 10, "number of iterations to show")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLifeServiceClient(conn)

	ctx := context.Background()

	r, err := c.GetRandomState(ctx, &pb.GetRandomStateRequest{
		Seed: *seed,
		Rows: *rows,
		Cols: *cols,
	})
	if err != nil {
		log.Fatalf("could not get random state: %v", err)
	}
	// log.Printf("response: {%+v}", r)
	gs, err := life.FromProto(r.GetState())
	if err != nil {
		log.Fatalf("invalid board: %v", err)
	}
	log.Printf("Board\n%s", gs.String())
	for i := range *iters {
		r, err := c.GetNextState(ctx, &pb.GetNextStateRequest{CurrentState: gs.ToProto()})
		if err != nil {
			log.Fatalf("could not get next state: %v", err)
		}
		gs, err = life.FromProto(r.GetNewState())
		log.Printf("[%d/%d] Board\n%s", i+1, *iters, gs.String())
	}
}
