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
		playerTypes   []pal.Type
		opponent      pal.Pal
		opponentSpeed int
		opponentTypes []pal.Type
		wantResult    Result
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
		{
			name:          "Fire vs Grass",
			player:        weakPal,
			playerSpeed:   50,
			playerTypes:   []pal.Type{pal.TypeFire},
			opponent:      weakPal,
			opponentSpeed: 50,
			opponentTypes: []pal.Type{pal.TypeGrass},
			wantResult:    ResultWin,
		},
		{
			name:          "Grass vs Fire",
			player:        weakPal,
			playerSpeed:   50,
			playerTypes:   []pal.Type{pal.TypeGrass},
			opponent:      weakPal,
			opponentSpeed: 50,
			opponentTypes: []pal.Type{pal.TypeFire},
			wantResult:    ResultLoose,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			player := test.player
			opponent := test.opponent
			if test.playerSpeed != 0 {
				player.Speed = test.playerSpeed
			}
			if test.opponentSpeed != 0 {
				opponent.Speed = test.opponentSpeed
			}
			if len(test.playerTypes) > 0 {
				player.Types = test.playerTypes
			}
			if len(test.opponentTypes) > 0 {
				opponent.Types = test.opponentTypes
			}
			gotResult := battler.Duel(player, opponent)
			require.Equal(t, test.wantResult, gotResult)
		})
	}
}
