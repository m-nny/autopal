package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"minmax.uk/autopal/pkg/life"
	"minmax.uk/autopal/pkg/life/client_tui"
	pb "minmax.uk/autopal/proto/life"
)

var (
	addr    = flag.String("addr", "localhost:50001", "address to connect to")
	timeout = flag.Duration("timeout", time.Second, "rpc timeout")

	rows = flag.Int64("rows", 10, "height of the board")
	cols = flag.Int64("cols", 20, "width of the board")
	seed = flag.Int64("seed", 42, "seed for init state")

	iters = flag.Int64("iters", 10, "number of iterations to show")

	tickDuration = flag.Duration("tick", time.Second, "tick duration")
)

func getRequest() *pb.PlayRandomGameRequest {
	until_stable := *iters == -1
	return &pb.PlayRandomGameRequest{
		Seed:            *seed,
		Rows:            *rows,
		Cols:            *cols,
		Iterations:      *iters,
		UntilStabilizes: until_stable,
	}
}

func runSimple(c pb.LifeServiceClient) error {
	ctx := context.Background()

	gs_stream, err := c.PlayRandomGame(ctx, getRequest())
	if err != nil {
		return fmt.Errorf("could not get random state: %w", err)
	}
	for {
		r, err := gs_stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return fmt.Errorf("could not get item from stream: %w", err)
		}
		gs, err := life.FromProto(r.GetState())
		if err != nil {
			return fmt.Errorf("recevied invalid game state: %w", err)
		}
		i := r.GetIteration()
		log.Printf("[%d/%d] Board\n%s\n\n", i+1, *iters, gs.String())
	}
	return nil
}

func runTui(c pb.LifeServiceClient) error {
	ctx := context.Background()
	stream, err := c.PlayRandomGame(ctx, getRequest())
	if err != nil {
		return err
	}
	model := client_tui.NewModel(stream, *tickDuration, *iters)
	logf, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		return err
	}
	defer logf.Close()
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		log.Printf("err: %v", err)
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLifeServiceClient(conn)

	if err := runTui(c); err != nil {
		log.Printf("failed to run tui: %v", err)
	}
}
