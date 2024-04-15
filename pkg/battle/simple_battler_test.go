package battle

import (
	"testing"

	"github.com/stretchr/testify/require"
	"minmax.uk/autopal/pkg/pal"
	"minmax.uk/autopal/pkg/pal/paltest"
)

func Test_SimpleBattler_Duel(t *testing.T) {
	paltest.Prep(t)
	rand := paltest.Rand()
	battler := NewSimpleBattler(rand)
	weakPal, err := pal.NewPal(rand, "1")
	require.NoError(t, err)
	strongPal, err := pal.NewPal(rand, "108")
	require.NoError(t, err)
	testCases := []struct {
		name          string
		player        pal.Pal
		playerSpeed   int
		opponent      pal.Pal
		opponentSpeed int
		wantResult    Result
	}{
		{
			name:          "Strong vs Weak pal",
			player:        strongPal,
			playerSpeed:   50,
			opponent:      weakPal,
			opponentSpeed: 50,
			wantResult:    ResultWin,
		},
		{
			name:          "Weak vs Strong pal",
			player:        weakPal,
			playerSpeed:   50,
			opponent:      strongPal,
			opponentSpeed: 50,
			wantResult:    ResultLoose,
		},
		{
			name:          "Same pal",
			player:        weakPal,
			playerSpeed:   50,
			opponent:      weakPal,
			opponentSpeed: 50,
			wantResult:    ResultDraw,
		},
		{
			name:          "Same pal but faster",
			player:        weakPal,
			playerSpeed:   51,
			opponent:      weakPal,
			opponentSpeed: 50,
			wantResult:    ResultWin,
		},
		{
			name:          "Same pal but slower",
			player:        weakPal,
			playerSpeed:   50,
			opponent:      weakPal,
			opponentSpeed: 51,
			wantResult:    ResultLoose,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.player.Speed = test.playerSpeed
			test.opponent.Speed = test.opponentSpeed
			gotResult := battler.Duel(test.player, test.opponent)
			require.Equal(t, test.wantResult, gotResult)
		})
	}
}
