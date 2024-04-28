package admin

import (
	"minmax.uk/autopal/pkg/admin/rpc"
	"minmax.uk/autopal/pkg/brain"
)

func newCreateUserModel(b *brain.Brain, username string) submodelItem {
	model := rpc.New("Creating new user", func() (*brain.User, error) {
		user, err := b.CreateUser(username)
		return user, err
	})
	return submodelItem{"Create user", model}
}

func newGetUserInfoModel(b *brain.Brain, username string) submodelItem {
	model := rpc.New("Geting user", func() (*brain.User, error) {
		user, err := b.GetUser(username)
		return user, err
	})
	return submodelItem{"Get user", model}
}
