package pal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPal(t *testing.T) {
	testLoadPalBases(t)
	testCases := []struct {
		name    string
		id      PalBaseId
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
			got, err := NewPal(test.id)
			if test.wantErr {
				require.Error(err)
				return
			} else {
				require.NoError(err)
			}
			require.Equal(got.Id, test.id)
			require.NotEmpty(got.Name)
			require.GreaterOrEqual(got.Speed, SPEED_MIN)
			require.LessOrEqual(got.Speed, SPEED_MAX)
		})
	}
}
