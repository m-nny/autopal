package life

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	Underpopulated lipgloss.Style
	Fine           lipgloss.Style
	Overpopulated  lipgloss.Style
	Baby           lipgloss.Style
}

func DefaultStyles() (s Styles) {
	s.Underpopulated = lipgloss.NewStyle().
		Foreground(lipgloss.Color("202"))
	s.Fine = lipgloss.NewStyle()
	s.Overpopulated = lipgloss.NewStyle().
		Foreground(lipgloss.Color("196"))
	s.Baby = lipgloss.NewStyle().
		Foreground(lipgloss.Color("40"))
	return s
}

var _defaultStyle = DefaultStyles()

func (state *GameState) GetColoredBoard() string {
	if state == nil {
		return "<empty>"
	}
	n_board := state.NBoard()
	var lines []string
	for row := range state.Rows {
		var builder strings.Builder
		for col := range state.Cols {
			i := state.AbsPos(col, row)
			r := cell_to_char[state.cells[i]]
			style := _defaultStyle.Fine
			if state.cells[i] {
				if n_board[i] < 2 {
					style = _defaultStyle.Underpopulated
				} else if 2 <= n_board[i] && n_board[i] <= 3 {
					style = _defaultStyle.Baby
				} else if 4 <= n_board[i] {
					style = _defaultStyle.Overpopulated
				}
			} else {
				if n_board[i] == 3 {
					style = _defaultStyle.Baby
				}
			}
			builder.WriteString(style.Render(string(r)))
		}
		lines = append(lines, builder.String())
	}
	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}
