package client_tui

import (
	"errors"
	"fmt"
	"io"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"minmax.uk/autopal/pkg/life"
	pb "minmax.uk/autopal/proto/life"
)

type Model struct {
	stream       pb.LifeService_PlayRandomGameClient
	tickDuration time.Duration
	totalIters   int64

	state   *life.GameState
	n_board []int

	curIter  int64
	err      error
	finished bool
}

func NewModel(stream pb.LifeService_PlayRandomGameClient, tickDuration time.Duration, totalIters int64) *Model {
	return &Model{
		stream:       stream,
		tickDuration: tickDuration,
		totalIters:   totalIters,

		state:   nil,
		n_board: nil,

		curIter:  0,
		err:      nil,
		finished: false,
	}
}

var _ tea.Model = (*Model)(nil)

type newStateMsg = *pb.PlayRandomGameResponse

type tickMsg time.Time

type errMsg error

type gameOverMsg int

func (m *Model) doTick() tea.Cmd {
	return tea.Tick(m.tickDuration, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(m.doTick(), m.getState())
}

func (m *Model) getState() tea.Cmd {
	return func() tea.Msg {
		if m.err != nil {
			return nil
		}
		if m.stream == nil {
			return errMsg(fmt.Errorf("stream is not initialized"))
		}
		resp, err := m.stream.Recv()
		if errors.Is(err, io.EOF) {
			return gameOverMsg(0)
		}
		if err != nil {
			return errMsg(err)
		}
		return newStateMsg(resp)
	}
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// log.Printf("Model.Update() msg: %#v", msg)
	switch msg := msg.(type) {
	case errMsg:
		m.err = msg
		return m, nil
	case gameOverMsg:
		m.finished = true
		return m, nil
	case tickMsg:
		if !m.finished {
			return m, tea.Batch(m.getState(), m.doTick())
		}
	case newStateMsg:
		gs, err := life.FromProto(msg.GetState())
		if err != nil {
			m.err = err
		} else {
			m.state = gs
			m.n_board = m.state.NBoard()
		}
		m.curIter = msg.GetIteration()
		m.totalIters = msg.GetTotalIterations()

		return m, nil
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *Model) View() string {
	var lines []string
	header := fmt.Sprintf("%3d of %3d", m.curIter, m.totalIters)
	if m.finished {
		header += " done"
	}
	lines = append(lines, header)
	if m.err != nil {
		lines = append(lines, fmt.Sprintf("Got error: %v", m.err.Error()))
	} else if m.state != nil {
		lines = append(lines, m.state.GetColoredBoard())
	} else {
		lines = append(lines, "Waiting...")
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}
