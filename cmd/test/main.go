package main

import (
	"flag"
	"log"

	"minmax.uk/autopal/pkg/brain"
)

var dbDsn = flag.String("dsn", "file:./data/turso.db", "filepath to db")
var username = flag.String("username", "minmax", "username")

func main() {
	flag.Parse()
	b, err := brain.New(*dbDsn)
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()

	user, err := b.CreateUser(*username)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("user: %+v", user)

	log.Print("done.")
}
