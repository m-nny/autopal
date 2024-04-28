package server

import (
	"context"

	pb "minmax.uk/autopal/proto"
)

func (s *Server) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	user, err := s.b.GetUser(req.GetUsername())
	if err != nil {
		return nil, err
	}
	return &pb.GetUserInfoResponse{
		UserInfo: &pb.UserInfo{
			Username: user.Username,
		},
	}, nil
}
