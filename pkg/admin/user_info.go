package admin

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	pb "minmax.uk/autopal/proto"
)

var _ tea.Model = (*UserInfoModel)(nil)

type gotUserInfoMsg *pb.GetUserInfoResponse

type UserInfoModel struct {
	c        pb.MainServiceClient
	username string
	userInfo *pb.UserInfo
	err      error
}

func NewUserInfoModel(c pb.MainServiceClient, username string) *UserInfoModel {
	return &UserInfoModel{c, username, nil, nil}
}

func (m *UserInfoModel) Init() tea.Cmd {
	return func() tea.Msg {
		// TODO: figure out how to get proper context here
		ctx := context.TODO()
		res, err := m.c.GetUserInfo(ctx, &pb.GetUserInfoRequest{Username: m.username})
		if err != nil {
			return errMsg(err)
		}
		return gotUserInfoMsg(res)
	}
}
func (m *UserInfoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case gotUserInfoMsg:
		m.userInfo = msg.UserInfo
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
func (m *UserInfoModel) View() string {
	if m.err != nil {
		return fmt.Sprintf("\nGot error: %v\n\n", m.err)
	}
	s := fmt.Sprint("Requesting UserInfo ...\n")
	if m.userInfo != nil {
		s += fmt.Sprintf("Got UserInfo: {%+v}\n", m.userInfo)
	}
	return s
}
