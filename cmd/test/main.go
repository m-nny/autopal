package main

import (
	"flag"
	"log"

	"minmax.uk/autopal/pkg/pal"
)

func main() {
	flag.Parse()
	if err := pal.LoadPalBases(); err != nil {
		log.Fatalf("could not load PalBases: %v", err)
	}
	for item := range pal.AllPalBases() {
		log.Printf("item: %+v", item)
	}
}
