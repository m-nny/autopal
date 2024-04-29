package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"minmax.uk/autopal/pkg/life/boards"
	"minmax.uk/autopal/pkg/life/simple_tui"
)

func main() {
	logf, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}
	defer logf.Close()
	tickDuration := 100 * time.Millisecond
	p := tea.NewProgram(simple_tui.NewModel(boards.Pulsar(), tickDuration))
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
