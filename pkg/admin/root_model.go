package admin

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"minmax.uk/autopal/pkg/admin/common"
	"minmax.uk/autopal/pkg/brain"
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 0)
)

var _ list.DefaultItem = submodelItem{}

type submodelItem struct {
	title string
	model tea.Model
}

func (i submodelItem) Description() string {
	return i.title
}
func (i submodelItem) Title() string {
	return i.title
}
func (i submodelItem) FilterValue() string {
	return i.title
}

var _ tea.Model = (*rootModel)(nil)

type rootModel struct {
	list           list.Model
	activeSubmodel tea.Model
}

func NewRootModel(b *brain.Brain, username string) *rootModel {
	items := []list.Item{
		newCreateUserModel(b, username),
		newGetUserInfoModel(b, username),
	}
	ld := list.NewDefaultDelegate()
	ld.ShowDescription = false
	l := list.New(items, ld, 0, 0)
	l.Title = "What do you want to do?"
	return &rootModel{
		list: l,
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
		case "enter":
			item, ok := m.list.SelectedItem().(submodelItem)
			if ok {
				m.activeSubmodel = item.model
				return m, m.activeSubmodel.Init()
			}
		}
	case tea.WindowSizeMsg:
		w, h := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-w, msg.Height-h)
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *rootModel) View() string {
	if m.activeSubmodel != nil {
		return m.activeSubmodel.View()
	}
	return docStyle.Render(m.list.View())
}
