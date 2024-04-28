package rpc

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"minmax.uk/autopal/pkg/admin/common"
)

type errMsg error
type gotResult[R any] struct{ result R }

type Getter[R any] func() (R, error)

var _ tea.Model = (*Model[any])(nil)

type Model[R any] struct {
	spinner spinner.Model
	Title   string
	fn      Getter[R]
	done    bool
	result  R
	err     error

	Styles Styles
}

func New[R any](title string, fn Getter[R]) *Model[R] {
	styles := DefaultStyles()
	return &Model[R]{
		Title: title,
		spinner: spinner.New(
			spinner.WithSpinner(spinner.Ellipsis),
			spinner.WithStyle(styles.Spinner),
		),
		fn:     fn,
		Styles: styles,
	}
}

func (m *Model[R]) Init() tea.Cmd {
	return tea.Batch(func() tea.Msg {
		res, err := m.fn()
		if err != nil {
			return errMsg(err)
		}
		return gotResult[R]{res}
	}, m.spinner.Tick)
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
	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m *Model[R]) View() string {
	sections := []string{m.Styles.Title.Render(m.Title)}
	if !m.done {
		sections = append(sections, m.spinner.View())
	} else {
		if m.err != nil {
			sections = append(sections, m.Styles.ErrorMessage.Render("Got error:"))
			sections = append(sections, fmt.Sprintf("%#v", m.err))
		} else {
			sections = append(sections, m.Styles.OkMessage.Render("Done:"))
			sections = append(sections, fmt.Sprintf("%#v", m.result))
		}
	}
	container := lipgloss.JoinVertical(lipgloss.Left, sections...)
	return m.Styles.Container.Render(container)
}
