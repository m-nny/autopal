package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"minmax.uk/autopal/pkg/admin"
	"minmax.uk/autopal/pkg/brain"
	pb "minmax.uk/autopal/proto"
)

var (
	addr     = flag.String("addr", "localhost:50000", "address to connect to")
	username = flag.String("name", "minmax", "name to report")
	timeout  = flag.Duration("timeout", time.Second, "rpc timeout")
	dbDsn    = flag.String("dsn", "file:./data/turso.db", "filepath to db")
)

func getClient(addr string) (pb.MainServiceClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}
	c := pb.NewMainServiceClient(conn)
	return c, nil
}

func main() {
	flag.Parse()
	b, err := brain.New(*dbDsn)
	if err != nil {
		log.Fatal(err)
	}
	// p := tea.NewProgram(admin.NewUserInfoModel(b, *username))
	// p := tea.NewProgram(admin.NewCreateUserModel(b, *username))
	logf, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}
	defer logf.Close()
	p := tea.NewProgram(admin.NewRootModel(b, *username))
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
