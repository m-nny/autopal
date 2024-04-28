package admin

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"minmax.uk/autopal/pkg/brain"
)

var _ tea.Model = (*CreateUserModel)(nil)

type gotCreateUserMsg *brain.User

type CreateUserModel struct {
	b        *brain.Brain
	username string
	user     *brain.User
	err      error
}

func NewCreateUserModel(b *brain.Brain, username string) *CreateUserModel {
	return &CreateUserModel{b, username, nil, nil}
}

func (m *CreateUserModel) Init() tea.Cmd {
	return func() tea.Msg {
		res, err := m.b.CreateUser(m.username)
		if err != nil {
			return errMsg(err)
		}
		return gotCreateUserMsg(res)
	}
}
func (m *CreateUserModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case gotCreateUserMsg:
		m.user = msg
		return m, tea.Quit
	case errMsg:
		m.err = msg
		return m, tea.Quit
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}
func (m *CreateUserModel) View() string {
	if m.err != nil {
		return fmt.Sprintf("\nGot error: %v\n\n", m.err)
	}
	s := fmt.Sprint("creating user ...\n")
	if m.user != nil {
		s += fmt.Sprintf("Got user: {%+v}\n", m.user)
	}
	return s
}
