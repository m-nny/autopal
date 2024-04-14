package pal

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadPalBases(t *testing.T) {
	updateDataDirFlag()
	defer resetDataDirFlag()
	err := LoadPalBases()
	if assert.NoError(t, err) {
		assert.True(t, len(_palBases) > 0)
		for item := range _palBases.All() {
			assert.NotEmpty(t, item.Name)
			assert.NotEmpty(t, item.Id)
			assert.NotEmpty(t, item.Types)
			assert.True(t, len(item.Types) > 0)
		}
	}
}

func updateDataDirFlag() {
	flag.Set("data_dir", "../../data/")
}
func resetDataDirFlag() {
	flag.Set("data_dir", "./data/")
}
