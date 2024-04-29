package main

import (
	"context"
	"errors"
	"flag"
	"io"
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

	iters = flag.Int64("iters", 10, "number of iterations to show")
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

	until_stable := *iters == -1
	gs_stream, err := c.PlayRandomGame(ctx, &pb.PlayRandomGameRequest{
		Seed:            *seed,
		Rows:            *rows,
		Cols:            *cols,
		Iterations:      *iters,
		UntilStabilizes: until_stable,
	})
	if err != nil {
		log.Fatalf("could not get random state: %v", err)
	}
	for {
		r, err := gs_stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatalf("could not get item from stream: %v", err)
		}
		gs, err := life.FromProto(r.GetState())
		if err != nil {
			log.Fatalf("recevied invalid game state: %v", err)
		}
		i := r.GetIteration()
		log.Printf("[%d/%d] Board\n%s\n\n", i+1, *iters, gs.String())
	}
}
