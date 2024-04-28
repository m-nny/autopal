package rpc

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"minmax.uk/autopal/pkg/admin/common"
)

var _ tea.Model = (*Model[any])(nil)

type errMsg error
type gotResult[R any] struct{ result R }

type Getter[R any] func() (R, error)

type Model[R any] struct {
	fn     Getter[R]
	done   bool
	result R
	err    error
}

func New[R any](fn Getter[R]) *Model[R] {
	return &Model[R]{fn: fn}
}

func (m *Model[R]) Init() tea.Cmd {
	return func() tea.Msg {
		res, err := m.fn()
		if err != nil {
			return errMsg(err)
		}
		return gotResult[R]{res}
	}
}

func (m *Model[R]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case gotResult[R]:
		m.result = msg.result
		m.done = true
		return m, nil
	case errMsg:
		m.err = msg
		m.done = true
		return m, nil
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter", "esc":
			if m.done {
				return m, common.CmdHome
			}
		}
	}
	return m, nil
}

func (m *Model[R]) View() string {
	if m.err != nil {
		return fmt.Sprintf("\nGot error: %v\n\n", m.err)
	}
	s := fmt.Sprint("creating user ...\n")
	if m.done {
		s += fmt.Sprintf("Got user: {%+v}\n", m.result)
	}
	return s
}
