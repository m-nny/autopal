package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"minmax.uk/autopal/pkg/admin"
)

func main() {
	p := tea.NewProgram(admin.NewModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
