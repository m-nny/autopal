package life

import (
	"fmt"
	"slices"
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
	Cols    int
	Rows    int
	cells   []bool
	n_board []int
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
	return &GameState{cols, rows, cells, nil}
}

func (s *GameState) String() string {
	var builder strings.Builder
	for row := range s.Rows {
		for col := range s.Cols {
			i := row*s.Cols + col
			cell := s.cells[i]
			builder.WriteRune(cell_to_char[cell])
		}
		if row+1 < s.Rows {
			builder.WriteRune('\n')
		}
	}
	return builder.String()
}

func (s *GameState) NBoard() []int {
	if s.n_board != nil {
		return s.n_board
	}
	n_board := make([]int, s.Cols*s.Rows, s.Cols*s.Rows)
	for row := range s.Rows {
		for col := range s.Cols {
			neighbours := 0
			for d_row := -1; d_row <= 1; d_row++ {
				for d_col := -1; d_col <= 1; d_col++ {
					if d_row == 0 && d_col == 0 {
						continue
					}
					neighbours += s.cntAt(col+d_col, row+d_row)
				}
			}
			i := s.AbsPos(col, row)
			n_board[i] = neighbours
		}
	}
	s.n_board = n_board
	return n_board
}

func (s *GameState) Next() *GameState {
	new_s := EmptyGame(s.Cols, s.Rows)
	n_board := s.NBoard()
	for i := range s.Rows * s.Cols {
		if s.cells[i] && 2 <= n_board[i] && n_board[i] <= 3 {
			// Any live cell with two or three live neighbors lives on to the next generation.
			new_s.cells[i] = true
		}
		if !s.cells[i] && n_board[i] == 3 {
			new_s.cells[i] = true
		}
	}
	return new_s
}

func (s *GameState) normPos(col, row int) (int, int) {
	col = (col + s.Cols) % s.Cols
	row = (row + s.Rows) % s.Rows
	return col, row
}

func (s *GameState) AbsPos(col, row int) int {
	i := row*s.Cols + col
	return i
}

func (s *GameState) cntAt(col, row int) int {
	i := s.AbsPos(s.normPos(col, row))
	if s.cells[i] {
		return 1
	}
	return 0
}

func FromCells(cols int, rows int, cells []bool) (*GameState, error) {
	gs := &GameState{
		Cols:  cols,
		Rows:  rows,
		cells: cells,
	}
	return gs, gs.Valid()
}

func FromProto(pgs *pb.GameState) (*GameState, error) {
	if pgs == nil {
		return nil, fmt.Errorf("pgs is nil")
	}
	gs := &GameState{
		Cols:  int(pgs.GetCols()),
		Rows:  int(pgs.GetRows()),
		cells: pgs.GetCells(),
	}
	return gs, gs.Valid()
}

func (gs *GameState) Valid() error {
	if gs.Cols <= 0 || gs.Rows <= 0 {
		return fmt.Errorf("both cols and rows shoule be positive: cols %d rows %d", gs.Cols, gs.Rows)
	}
	if len(gs.cells) != gs.Cols*gs.Rows {
		return fmt.Errorf("number of cells does not match cols*rows len(%d) != %d*%d", len(gs.cells), gs.Cols, gs.Rows)
	}
	return nil
}

func (s *GameState) ToProto() *pb.GameState {
	return &pb.GameState{
		Cols:  int64(s.Cols),
		Rows:  int64(s.Rows),
		Cells: s.cells,
	}
}

func (s *GameState) Equal(other *GameState) bool {
	if s.Cols != other.Cols || s.Rows != other.Rows {
		return false
	}
	return slices.Equal(s.cells, other.cells)
}
