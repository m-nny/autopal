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
		assert.NotEmpty(t, len(_palBases))
		for item := range _palBases.All() {
			assert.NotEmpty(t, item.Name)
			assert.NotEmpty(t, item.Id)
			assert.NotEmpty(t, item.Types)
			assert.NotEmpty(t, len(item.Types))
		}
	}
}

func updateDataDirFlag() {
	flag.Set("data_dir", "../../data/")
}
func resetDataDirFlag() {
	flag.Set("data_dir", "./data/")
}
