package server

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"minmax.uk/autopal/pkg/brain"
	"minmax.uk/autopal/pkg/brain/brain_testing"
	pb "minmax.uk/autopal/proto"
)

func Test_GetUserInfo_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	require := require.New(t)
	ctx := context.Background()
	b := brain_testing.NewTestBrain(t)
	client, closer := NewTestClient(t, ctx, b)
	defer closer()

	username := "test_username"
	_, err := client.GetUserInfo(ctx, &pb.GetUserInfoRequest{Username: username})
	require.Error(err)
	require.Equal(codes.NotFound, status.Code(err), "returned error code is not NotFound")

	_, err = b.CreateUser(username)
	require.NoError(err)

	want := &pb.GetUserInfoResponse{UserInfo: &pb.UserInfo{Username: username, Balance: brain.UserStartingBalance}}
	got, err := client.GetUserInfo(ctx, &pb.GetUserInfoRequest{Username: username})
	require.NoError(err)
	require.EqualExportedValues(want, got)
}
