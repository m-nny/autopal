package paltest

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/require"
	"minmax.uk/autopal/pkg/pal"
)

func Prep(t testing.TB) {
	require := require.New(t)
	flag.Set("data_dir", "../../data/")
	err := pal.LoadPalBases()
	require.NoError(err)
}
