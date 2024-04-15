package battle

import (
	"testing"

	"github.com/stretchr/testify/require"
	"minmax.uk/autopal/pkg/pal"
	"minmax.uk/autopal/pkg/pal/paltest"
)

func Test_SimpleBattler_Duel(t *testing.T) {
	paltest.Prep(t)
	battler := NewSimpleBattler()
	weakPal, err := pal.NewPal("1")
	require.NoError(t, err)
	strongPal, err := pal.NewPal("108")
	require.NoError(t, err)
	testCases := []struct {
		name       string
		player     pal.Pal
		opponent   pal.Pal
		wantResult Result
	}{
		{
			name:       "Strong vs Weak pal",
			player:     strongPal,
			opponent:   weakPal,
			wantResult: ResultWin,
		},
		{
			name:       "Weak vs Strong pal",
			player:     weakPal,
			opponent:   strongPal,
			wantResult: ResultLoose,
		},
		{
			name:       "Same pal",
			player:     weakPal,
			opponent:   weakPal,
			wantResult: ResultDraw,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			gotResult := battler.Duel(test.player, test.opponent)
			require.Equal(t, test.wantResult, gotResult)
		})
	}
}
