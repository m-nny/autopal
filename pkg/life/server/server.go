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
