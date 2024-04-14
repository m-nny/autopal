package pal

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadPalBases(t *testing.T) {
	require := require.New(t)
	testLoadPalBases(t)
	require.NotEmpty(len(_palBases))
	for item := range _palBases.All() {
		require.NotEmpty(item.Id, item)
		require.NotEmpty(item.Name)
		require.Positive(item.BaseHp, item)
		require.Positive(item.BaseAttack)
		require.Positive(item.BaseDefence)
		require.NotEmpty(item.Types)
	}
}

func testLoadPalBases(t testing.TB) {
	require := require.New(t)
	flag.Set("data_dir", "../../data/")
	err := LoadPalBases()
	require.NoError(err)
}
