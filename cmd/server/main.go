package main

import (
	"flag"
	"fmt"
	"log"

	"minmax.uk/autopal/pkg/brain"
	"minmax.uk/autopal/pkg/server"
)

var (
	port  = flag.Int("port", 50000, "server port")
	dbDsn = flag.String("dsn", "file:./data/turso.db", "filepath to db")
)

func main() {
	flag.Parse()
	addr := fmt.Sprintf(":%d", *port)
	b, err := brain.New(*dbDsn)
	if err != nil {
		log.Fatal(err)
	}
	s := server.NewServer(b)
	if err := s.Serve(addr); err != nil {
		log.Fatal(err)
	}
}
