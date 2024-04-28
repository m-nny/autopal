package brain_testing

import (
	"testing"

	"github.com/stretchr/testify/require"
	"minmax.uk/autopal/pkg/brain"
)

func NewTestBrain(t testing.TB) *brain.Brain {
	dsn := ":memory:"
	b, err := brain.New(dsn)
	require.NoError(t, err)
	return b
}
