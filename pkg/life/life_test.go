package life

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_EmptyGame(t *testing.T) {
	require := require.New(t)
	rows := 3
	cols := 4
	gotState := EmptyGame(cols, rows)
	wantState := &GameState{
		cols: cols, rows: rows,
		cells: []bool{false, false, false, false, false, false, false, false, false, false, false, false},
	}
	require.Equal(wantState, gotState)

	wantStr := "....\n" + "....\n" + "...."
	gotStr := gotState.String()
	require.Equal(wantStr, gotStr)
}

func Test_FromString(t *testing.T) {
	testCases := []struct {
		name      string
		rows      int
		cols      int
		str       string
		wantErr   bool
		wantCells []bool
	}{
		{
			name: "empty board with dots",
			rows: 3,
			cols: 4,
			str: `
			....
			....
			....
`,
			wantCells: []bool{false, false, false, false, false, false, false, false, false, false, false, false},
		},
		{
			name: "empty board with dots",
			rows: 3,
			cols: 4,
			str: `
			#...
			.#..
			..#.
`,
			wantCells: []bool{true, false, false, false, false, true, false, false, false, false, true, false},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			require := require.New(t)
			gotState, err := FromString(test.cols, test.rows, test.str)
			if test.wantErr {
				require.Error(err)
				return
			}
			require.NoError(err)
			wantState := &GameState{cols: test.cols, rows: test.rows, cells: test.wantCells}
			require.Equal(wantState, gotState)
		})
	}
}

func Test_Next_OnOsclillators(t *testing.T) {
	oscilators := []struct {
		name        string
		rows        int
		cols        int
		init_state  string
		want_boards []string
	}{
		{
			name: "Blinker",
			rows: 5,
			cols: 5,
			init_state: `
				.....
				..#..
				..#..
				..#..
				.....
			`,
			want_boards: []string{`
				.....
				.....
				.###.
				.....
				.....
			`, `
				.....
				..#..
				..#..
				..#..
				.....
			`},
		},
	}
	for _, test := range oscilators {
		t.Run(test.name, func(t *testing.T) {
			require := require.New(t)
			gotState, err := FromString(test.cols, test.rows, test.init_state)
			require.NoError(err)
			for i, want_cell := range test.want_boards {
				gotState = gotState.Next()
				requireState(t, want_cell, gotState, "iteration #%d does not match", i+1)
			}
		})
	}
}

func requireState(t testing.TB, wantBoard string, gotState *GameState, msgAndArgs ...any) {
	wantState, err := FromString(gotState.cols, gotState.rows, wantBoard)
	require.NoError(t, err)
	require.Equal(t, wantState.cells, gotState.cells, msgAndArgs...)
}
