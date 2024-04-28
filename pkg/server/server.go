package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"minmax.uk/autopal/pkg/brain"
	pb "minmax.uk/autopal/proto"
)

type Server struct {
	b *brain.Brain
	pb.UnimplementedMainServiceServer
}

func NewServer(b *brain.Brain) *Server {
	return &Server{b: b}
}

func (s *Server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := req.GetName()
	log.Printf("[SayHello] received: {%+v}", req)
	msg := fmt.Sprintf("Hello %s", name)
	return &pb.HelloResponse{Message: msg}, nil
}

func (s *Server) Serve(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	grpcS := grpc.NewServer()
	pb.RegisterMainServiceServer(grpcS, s)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcS.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}
