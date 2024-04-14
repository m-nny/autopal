package pal

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadPalBases(t *testing.T) {
	require := require.New(t)
	updateDataDirFlag()
	defer resetDataDirFlag()
	err := LoadPalBases()
	require.NoError(err)
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

func updateDataDirFlag() {
	flag.Set("data_dir", "../../data/")
}
func resetDataDirFlag() {
	flag.Set("data_dir", "./data/")
}
