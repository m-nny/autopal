package rpc

import "github.com/charmbracelet/lipgloss"

type Styles struct {
	TitleBar     lipgloss.Style
	Title        lipgloss.Style
	Spinner      lipgloss.Style
	ErrorMessage lipgloss.Style
	OkMessage    lipgloss.Style
	Container    lipgloss.Style
}

func DefaultStyles() (s Styles) {
	s.TitleBar = lipgloss.NewStyle().
		Padding(0, 0, 1, 2)

	s.Title = lipgloss.NewStyle().
		Background(lipgloss.Color("62")).
		Foreground(lipgloss.Color("230")).
		Padding(0, 1)

	s.Spinner = lipgloss.NewStyle().
		Foreground(lipgloss.Color("205"))

	s.ErrorMessage = lipgloss.NewStyle().
		Foreground(lipgloss.Color("205"))

	s.OkMessage = lipgloss.NewStyle().
		Foreground(lipgloss.Color("149"))

	s.Container = lipgloss.NewStyle().
		Padding(0, 0, 1, 2)
	return s
}
