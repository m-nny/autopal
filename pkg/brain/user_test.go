package brain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CreateGetUser(t *testing.T) {
	require := require.New(t)
	b := NewTestBrain(t)
	username := "test_username"
	want := &User{username}

	_, err := b.GetUser(username)
	require.ErrorIs(err, ErrNotFound)

	got1, err := b.CreateUser(username)
	require.NoError(err)
	require.Equal(want, got1)

	got2, err := b.GetUser(username)
	require.NoError(err, ErrNotFound)
	require.Equal(want, got2)

	_, err = b.CreateUser(username)
	require.ErrorIs(err, ErrAlreadyExists)
}
