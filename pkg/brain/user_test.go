package brain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"minmax.uk/autopal/pkg/brain"
	brain_t "minmax.uk/autopal/pkg/brain/brain_testing"
)

func Test_CreateGetUser(t *testing.T) {
	require := require.New(t)
	b := brain_t.NewTestBrain(t)
	username := "test_username"
	want := &brain.User{username, 10}

	_, err := b.GetUser(username)
	require.ErrorIs(err, brain.ErrNotFound)

	got1, err := b.CreateUser(username)
	require.NoError(err)
	require.Equal(want, got1)

	got2, err := b.GetUser(username)
	require.NoError(err, brain.ErrNotFound)
	require.Equal(want, got2)

	_, err = b.CreateUser(username)
	require.ErrorIs(err, brain.ErrAlreadyExists)
}
