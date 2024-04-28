package admin

import (
	tea "github.com/charmbracelet/bubbletea"
	"minmax.uk/autopal/pkg/admin/rpc"
	"minmax.uk/autopal/pkg/brain"
)

func NewCreateUserModel(b *brain.Brain, username string) tea.Model {
	return rpc.New(func() (*brain.User, error) {
		user, err := b.CreateUser(username)
		return user, err
	})
}

func NewGetUserInfoModel(b *brain.Brain, username string) tea.Model {
	return rpc.New(func() (*brain.User, error) {
		user, err := b.GetUser(username)
		return user, err
	})
}
