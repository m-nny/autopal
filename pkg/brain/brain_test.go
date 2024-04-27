package brain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func NewTestBrain(t testing.TB) *Brain {
	require := require.New(t)
	dsn := ":memory:"
	b, err := New(dsn)
	require.NoError(err)
	return b
}
