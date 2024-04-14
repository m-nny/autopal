package main

import (
	"flag"
	"fmt"
	"log"

	"minmax.uk/autopal/pkg/pal"
)

func main() {
	flag.Parse()
	if err := pal.LoadPalBases(); err != nil {
		log.Fatalf("could not load PalBases: %v", err)
	}
	for item := range pal.AllPalBases() {
		fmt.Printf("%s\n\t%+v\n", item.String(), item)
	}
}
