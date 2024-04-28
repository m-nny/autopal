package server

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"minmax.uk/autopal/pkg/brain"
	pb "minmax.uk/autopal/proto"
)

func (s *Server) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	user, err := s.b.GetUser(req.GetUsername())
	if err != nil {
		return nil, toGrpcErr(err)
	}
	return &pb.GetUserInfoResponse{
		UserInfo: &pb.UserInfo{
			Username: user.Username,
		},
	}, nil
}

func toGrpcErr(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, brain.ErrNotFound) {
		return status.Error(codes.NotFound, err.Error())
	}
	if errors.Is(err, brain.ErrAlreadyExists) {
		return status.Error(codes.AlreadyExists, err.Error())
	}
	return err
}
