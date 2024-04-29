package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"minmax.uk/autopal/pkg/life"
	"minmax.uk/autopal/pkg/life/boards"
	pb "minmax.uk/autopal/proto/life"
)

const MAX_ITERS = 1000_000

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

	gs, err := boards.Rnadom(cols, rows, req.GetSeed())
	if err != nil {
		return nil, err
	}

	return &pb.GetRandomStateResponse{State: gs.ToProto()}, nil
}

func (s *Server) GetNextState(ctx context.Context, req *pb.GetNextStateRequest) (*pb.GetNextStateResponse, error) {
	init_gs, err := life.FromProto(req.GetCurrentState())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	new_gs := init_gs.Next()
	return &pb.GetNextStateResponse{NewState: new_gs.ToProto()}, nil
}

func (s *Server) PlayRandomGame(req *pb.PlayRandomGameRequest, stream pb.LifeService_PlayRandomGameServer) error {
	rows, cols := req.GetRows(), req.GetCols()
	if rows <= 0 || cols <= 0 {
		return status.Error(codes.InvalidArgument, "both rows and cols should be > 0")
	}
	iters := req.GetIterations()
	if req.GetUntilStabilizes() {
		iters = MAX_ITERS
	} else if iters <= 0 {
		return status.Error(codes.InvalidArgument, "iters should be positive")
	}

	gs, err := boards.Rnadom(cols, rows, req.GetSeed())
	if err != nil {
		return err
	}

	stabilized := false
	for i := range iters {
		stream.Send(&pb.PlayRandomGameResponse{
			Iteration: i,
			State:     gs.ToProto(),
		})
		new_gs := gs.Next()
		if new_gs.Equal(gs) {
			stabilized = true
			break
		}
		gs = new_gs
	}
	if req.GetUntilStabilizes() && !stabilized {
		return fmt.Errorf("game did not stabilize in %d iterations", MAX_ITERS)
	}
	return nil
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
