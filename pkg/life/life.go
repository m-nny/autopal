package life

import (
	"fmt"
	"strings"

	pb "minmax.uk/autopal/proto/life"
)

var cell_to_char = map[bool]rune{
	false: '.',
	true:  'O',
}
var EMPTY_CELLS = "."
var FULL_CELLS = "O#"
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
		if i == cols*rows {
			return nil, fmt.Errorf("too much chars")
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

func (s *GameState) Next() *GameState {
	new_s := EmptyGame(s.cols, s.rows)
	for row := range s.rows {
		for col := range s.cols {
			neighbours := 0
			for d_row := -1; d_row <= 1; d_row++ {
				for d_col := -1; d_col <= 1; d_col++ {
					if d_row == 0 && d_col == 0 {
						continue
					}
					neighbours += s.cntAt(col+d_col, row+d_row)
				}
			}
			i := s.absPos(col, row)
			if s.cells[i] && 2 <= neighbours && neighbours <= 4 {
				// Any live cell with two or three live neighbors lives on to the next generation.
				new_s.cells[i] = true
			}
			if !s.cells[i] && neighbours == 3 {
				new_s.cells[i] = true
			}
		}
	}
	return new_s
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

func FromCells(cols int, rows int, cells []bool) (*GameState, error) {
	gs := &GameState{
		cols:  cols,
		rows:  rows,
		cells: cells,
	}
	return gs, gs.Valid()
}

func FromProto(pgs *pb.GameState) (*GameState, error) {
	if pgs == nil {
		return nil, fmt.Errorf("pgs is nil")
	}
	gs := &GameState{
		cols:  int(pgs.GetCols()),
		rows:  int(pgs.GetRows()),
		cells: pgs.GetCells(),
	}
	return gs, gs.Valid()
}

func (gs *GameState) Valid() error {
	if gs.cols <= 0 || gs.rows <= 0 {
		return fmt.Errorf("both cols and rows shoule be positive: cols %d rows %d", gs.cols, gs.rows)
	}
	if len(gs.cells) != gs.cols*gs.rows {
		return fmt.Errorf("number of cells does not match cols*rows len(%d) != %d*%d", len(gs.cells), gs.cols, gs.rows)
	}
	return nil
}

func (s *GameState) ToProto() *pb.GameState {
	return &pb.GameState{
		Cols:  int64(s.cols),
		Rows:  int64(s.rows),
		Cells: s.cells,
	}
}
