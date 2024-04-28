package admin

import (
	tea "github.com/charmbracelet/bubbletea"
	"minmax.uk/autopal/pkg/admin/common"
	"minmax.uk/autopal/pkg/brain"
)

var _ tea.Model = (*rootModel)(nil)

type rootModel struct {
	submodels      []tea.Model
	submodelNames  []string
	activeSubmodel tea.Model
	cursor         int
}

func NewRootModel(b *brain.Brain, username string) *rootModel {
	return &rootModel{
		submodels:     []tea.Model{NewGetUserInfoModel(b, username), NewCreateUserModel(b, username)},
		submodelNames: []string{"Get user", "Create user"},
	}
}

func (m *rootModel) Init() tea.Cmd {
	return nil
}

func (m *rootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.activeSubmodel != nil {
		if _, ok := msg.(common.MsgHome); ok {
			m.activeSubmodel = nil
			return m, nil
		}
		var cmd tea.Cmd
		m.activeSubmodel, cmd = m.activeSubmodel.Update(msg)
		return m, cmd
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor+1 < len(m.submodelNames) {
				m.cursor++
			}
		case "enter":
			m.activeSubmodel = m.submodels[m.cursor]
			return m, m.activeSubmodel.Init()
		}
	}
	return m, nil
}

func (m *rootModel) View() string {
	if m.activeSubmodel != nil {
		return m.activeSubmodel.View()
	}
	s := "What do you want to do?\n"
	for i, choice := range m.submodelNames {
		line := " "
		if i == m.cursor {
			line = ">"
		}
		line += " " + choice + "\n"
		s += line
	}
	s += "\nPress ender to select.\n"
	s += "\nPress q to quit.\n"
	return s
}
