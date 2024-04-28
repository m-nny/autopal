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
	cells []bool
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
		if strings.ContainsRune(EMPTY_CELLS, rune) {
			g.cells[i] = false
		} else if strings.ContainsRune(FULL_CELLS, rune) {
			g.cells[i] = true
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

func MustFromString(cols, rows int, str string) *GameState {
	g, err := FromString(cols, rows, str)
	if err != nil {
		panic(fmt.Sprintf("FromString() errored: %v", err))
	}
	return g
}

func EmptyGame(cols int, rows int) *GameState {
	cells := make([]bool, cols*rows, cols*rows)
	return &GameState{cols, rows, cells}
}

func (s *GameState) String() string {
	var builder strings.Builder
	for row := range s.rows {
		for col := range s.cols {
			i := row*s.cols + col
			cell := s.cells[i]
			builder.WriteRune(cell_to_char[cell])
		}
		if row+1 < s.rows {
			builder.WriteRune('\n')
		}
	}
	return builder.String()
}

func (s *GameState) normPos(col, row int) (int, int) {
	col = (col + s.cols) % s.cols
	row = (row + s.rows) % s.rows
	return col, row
}

func (s *GameState) absPos(col, row int) int {
	i := row*s.cols + col
	return i
}

func (s *GameState) cntAt(col, row int) int {
	i := s.absPos(s.normPos(col, row))
	if s.cells[i] {
		return 1
	}
	return 0
}
