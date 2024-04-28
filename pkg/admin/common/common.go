package common

import (
	tea "github.com/charmbracelet/bubbletea"
)

type MsgHome struct{ from string }

func CmdHome() tea.Msg {
	return MsgHome{from: "somewhere"}
}
