package pal

import (
	"fmt"
	"math/rand"
)

// Pal  replresents particular Pal (instance of PalBase)
type Pal struct {
	PalBase
	Speed int
}

const SPEED_MIN = 1
const SPEED_MAX = 100

func NewPal(rand *rand.Rand, id PalBaseId) (Pal, error) {
	base, ok := _palBases.Find(id)
	if !ok {
		return Pal{}, fmt.Errorf("pal with id %s not found", id)
	}
	return Pal{
		PalBase: base,
		Speed:   normRand(rand, SPEED_MIN, SPEED_MAX),
	}, nil
}

func normRand(rand *rand.Rand, min, max int) int {
	val := int(rand.NormFloat64())
	if val < min {
		val = min
	}
	if max < val {
		val = max
	}
	return val
}
