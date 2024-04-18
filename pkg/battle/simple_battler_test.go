package battle

import (
	"fmt"
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
			player:        strongPal,
			playerSpeed:   50,
			playerTypes:   []pal.Type{pal.TypeFire},
			opponent:      strongPal,
			opponentSpeed: 50,
			opponentTypes: []pal.Type{pal.TypeGrass},
			wantResult:    ResultWin,
		},
		{
			name:          "Grass vs Fire",
			player:        strongPal,
			playerSpeed:   50,
			playerTypes:   []pal.Type{pal.TypeGrass},
			opponent:      strongPal,
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

func Test_Attack(t *testing.T) {
	paltest.Prep(t)
	rand := paltest.Rand()
	testStats := []struct {
		baseAttack  int
		baseDefence int
		wantDmg     int
	}{
		// Neutral vs Neutral
		{
			baseAttack:  100,
			baseDefence: 100,
			wantDmg:     100,
		},
		{
			baseAttack:  100,
			baseDefence: 80,
			wantDmg:     120,
		},
		{
			baseAttack:  120,
			baseDefence: 100,
			wantDmg:     120,
		},
		{
			baseAttack:  100,
			baseDefence: 120,
			wantDmg:     80,
		},
		{
			baseAttack:  120,
			baseDefence: 120,
			wantDmg:     100,
		},
	}
	testTypes := []struct {
		player        []pal.Type
		opponent      []pal.Type
		wantScaleUp   int
		wantScaleDown int
	}{
		{
			player:        []pal.Type{pal.TypeNeutral},
			opponent:      []pal.Type{pal.TypeNeutral},
			wantScaleUp:   1,
			wantScaleDown: 1,
		},
		{
			player:        []pal.Type{pal.TypeDark},
			opponent:      []pal.Type{pal.TypeNeutral},
			wantScaleUp:   2,
			wantScaleDown: 1,
		},
		{
			player:        []pal.Type{pal.TypeNeutral},
			opponent:      []pal.Type{pal.TypeDark},
			wantScaleUp:   1,
			wantScaleDown: 2,
		},
		{
			player:        []pal.Type{pal.TypeFire},
			opponent:      []pal.Type{pal.TypeGrass, pal.TypeIce},
			wantScaleUp:   4,
			wantScaleDown: 1,
		},
	}
	for _, testType := range testTypes {
		for _, testStat := range testStats {
			test_name := fmt.Sprintf("%03d_%03d_%v_%v", testStat.baseAttack, testStat.baseDefence, testType.player, testType.opponent)
			t.Run(test_name, func(t *testing.T) {
				require := require.New(t)
				baseBal, err := pal.NewPal(rand, "78")
				require.NoError(err)

				baseBal.BaseAttack = testStat.baseAttack
				baseBal.BaseDefence = testStat.baseDefence

				player := newBattlePal(baseBal)
				if len(testType.player) > 0 {
					player.Types = testType.player
				}
				opponent := newBattlePal(baseBal)
				if len(testType.opponent) > 0 {
					opponent.Types = testType.opponent
				}
				oldHp := opponent.currentHp
				newHp := player.Attack(opponent)
				gotDmg := oldHp - newHp
				wantDmt := testStat.wantDmg * testType.wantScaleUp / testType.wantScaleDown
				// fmt.Printf("%s: %d\n", test_name, wantDmt)
				require.Equal(wantDmt, gotDmg)
			})
		}
	}
}
