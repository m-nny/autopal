package boards

import (
	"fmt"
	"math/rand"

	"minmax.uk/autopal/pkg/life"
)

func Random(cols, rows, seed int64) (*life.GameState, error) {
	if rows <= 0 || cols <= 0 {
		return nil, fmt.Errorf("both rows and cols should be > 0")
	}

	// TODO: add options for randomness
	rand := rand.New(rand.NewSource(seed))

	// We can do it more efficiently by generating whole chunks (int64) and using it
	cells := make([]bool, rows*cols, rows*cols)
	for i := range rows * cols {
		cells[i] = getRandBool(rand)
	}
	gs, err := life.FromCells(int(cols), int(rows), cells)
	if err != nil {
		return nil, err
	}
	return gs, nil
}

func getRandBool(rand *rand.Rand) bool {
	max_n := 100
	return (rand.Intn(max_n) & 1) == 1
}
