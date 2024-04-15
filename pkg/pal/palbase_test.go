package pal_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"minmax.uk/autopal/pkg/pal"
	"minmax.uk/autopal/pkg/pal/paltest"
)

func TestLoadPalBases(t *testing.T) {
	require := require.New(t)
	paltest.Prep(t)
	atLeastOne := false
	for item := range pal.AllPalBases() {
		atLeastOne = true
		require.NotEmpty(item.Id, item)
		require.NotEmpty(item.Name)
		require.Positive(item.BaseHp, item)
		require.Positive(item.BaseAttack)
		require.Positive(item.BaseDefence)
		require.NotEmpty(item.Types)
	}
	require.True(atLeastOne)
}
