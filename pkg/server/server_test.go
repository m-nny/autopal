package server

import (
	"context"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"minmax.uk/autopal/pkg/brain"
	pb "minmax.uk/autopal/proto"
)

func NewTestClient(t testing.TB, ctx context.Context, b *brain.Brain) (pb.MainServiceClient, func()) {
	bufferSize := 101024 * 1024
	lis := bufconn.Listen(bufferSize)

	baseServer := grpc.NewServer()
	pb.RegisterMainServiceServer(baseServer, NewServer(b))
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			t.Fatalf("error serving testing server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) { return lis.Dial() }),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("error connecting to test server: %v", err)
	}
	closer := func() {
		if err := lis.Close(); err != nil {
			t.Fatalf("error closing listetner: %v", err)
		}
		baseServer.Stop()
	}

	client := pb.NewMainServiceClient(conn)
	return client, closer
}
