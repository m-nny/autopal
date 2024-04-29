package simple_tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"minmax.uk/autopal/pkg/life"
)

type Model struct {
	state        *life.GameState
	tickDuration time.Duration
}

func NewModel(state *life.GameState, tickDuration time.Duration) *Model {
	return &Model{state, tickDuration}
}

var _ tea.Model = (*Model)(nil)

type TickMsg time.Time

func (m *Model) doTick() tea.Cmd {
	return tea.Tick(m.tickDuration, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}
func (m *Model) Init() tea.Cmd {
	return m.doTick()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case TickMsg:
		m.state = m.state.Next()
		return m, m.doTick()
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}
func (m *Model) View() string {
	return m.state.String()
}
