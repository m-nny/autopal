package server

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"minmax.uk/autopal/pkg/life"
	pb "minmax.uk/autopal/proto/life"
)

type Server struct {
	pb.UnimplementedLifeServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetRandomState(ctx context.Context, req *pb.GetRandomStateRequest) (*pb.GetRandomStateResponse, error) {
	rows, cols := req.GetRows(), req.GetCols()
	if rows <= 0 || cols <= 0 {
		return nil, status.Error(codes.InvalidArgument, "both rows and cols should be > 0")
	}

	// TODO: add options for randomness
	rand := rand.New(rand.NewSource(req.Seed))

	// We can do it more efficiently by generating whole chunks (int64) and using it
	cells := make([]bool, rows*cols, rows*cols)
	for i := range rows * cols {
		cells[i] = getRandBool(rand)
	}
	gs, err := life.FromCells(int(cols), int(rows), cells)
	if err != nil {
		return nil, err
	}
	return &pb.GetRandomStateResponse{State: gs.ToProto()}, nil
}

func (s *Server) Serve(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	grpcS := grpc.NewServer()
	pb.RegisterLifeServiceServer(grpcS, s)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcS.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}

func getRandBool(rand *rand.Rand) bool {
	max_n := 100
	return (rand.Intn(max_n) & 1) == 1

}
