package pal_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"minmax.uk/autopal/pkg/pal"
	"minmax.uk/autopal/pkg/pal/paltest"
)

func TestNewPal(t *testing.T) {
	paltest.Prep(t)
	testCases := []struct {
		name    string
		id      pal.PalBaseId
		wantErr bool
	}{
		{
			name:    "foxparks",
			id:      "5",
			wantErr: false,
		},
		{
			name:    "incorrect id",
			id:      "-5",
			wantErr: true,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			require := require.New(t)
			got, err := pal.NewPal(test.id)
			if test.wantErr {
				require.Error(err)
				return
			} else {
				require.NoError(err)
			}
			require.Equal(got.Id, test.id)
			require.NotEmpty(got.Name)
			require.GreaterOrEqual(got.Speed, pal.SPEED_MIN)
			require.LessOrEqual(got.Speed, pal.SPEED_MAX)
		})
	}
}
