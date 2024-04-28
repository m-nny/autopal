package life

import (
	"fmt"
	"strings"
)

var cell_to_char = map[bool]rune{
	false: '.',
	true:  '#',
}
var EMPTY_CELLS = "."
var FULL_CELLS = "#"
var SKIP_CELLS = "\n\t "

type GameState struct {
	cols  int
	rows  int
	cells [][]bool
}

func FromString(cols int, rows int, str string) (*GameState, error) {
	if len(str) < cols*rows {
		return nil, fmt.Errorf("not enough chars")
	}
	g := EmptyGame(cols, rows)
	i := 0
	for _, rune := range str {
		if strings.ContainsRune(SKIP_CELLS, rune) {
			continue
		}
		row := i / cols
		col := i % cols
		if strings.ContainsRune(EMPTY_CELLS, rune) {
			g.cells[row][col] = false
		} else if strings.ContainsRune(FULL_CELLS, rune) {
			g.cells[row][col] = true
		} else {
			return nil, fmt.Errorf("illegal rune: {%+v} {%d}", rune, rune)
		}
		i++
	}
	if i != cols*rows {
		return nil, fmt.Errorf("not enough chars")
	}
	return g, nil
}

func EmptyGame(cols int, rows int) *GameState {
	cells := make([][]bool, rows, rows)
	for row := range rows {
		cells[row] = make([]bool, cols, cols)
	}
	return &GameState{cols, rows, cells}
}

func (s *GameState) String() string {
	var builder strings.Builder
	for row := range s.rows {
		for col := range s.cols {
			cell := s.cells[row][col]
			builder.WriteRune(cell_to_char[cell])
		}
		if row+1 < s.rows {
			builder.WriteRune('\n')
		}
	}
	return builder.String()
}
