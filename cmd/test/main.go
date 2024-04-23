package main

import (
	"context"
	"flag"
	"log"

	"minmax.uk/autopal/pkg/brain"
)

var dbDsn = flag.String("dsn", "file:./data/turso.db", "filepath to db")
var username = flag.String("username", "minmax", "username")

func main() {
	ctx := context.Background()
	flag.Parse()
	b, err := brain.New(*dbDsn)
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()

	user, err := b.UpsertUser(ctx, *username)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("user: %+v", user)

	log.Print("done.")
}
