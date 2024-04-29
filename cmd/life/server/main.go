package main

import (
	"flag"
	"fmt"
	"log"

	"minmax.uk/autopal/pkg/life/server"
)

var (
	port = flag.Int("port", 50001, "server port")
)

func main() {
	flag.Parse()
	addr := fmt.Sprintf(":%d", *port)
	s := server.NewServer()
	if err := s.Serve(addr); err != nil {
		log.Fatal(err)
	}
}
