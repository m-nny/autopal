package battle

import (
	"log"
	"math/rand"

	"minmax.uk/autopal/pkg/pal"
)

type Result int

const (
	ResultUnknown Result = iota
	ResultWin
	ResultLoose
	ResultDraw
)

type SimpleBattler struct {
	rand *rand.Rand
}

func NewSimpleBattler(rand *rand.Rand) *SimpleBattler {
	return &SimpleBattler{rand: rand}
}

type battlePal struct {
	pal.Pal
	currentHp int
}

func newBattlePal(p pal.Pal) *battlePal {
	return &battlePal{
		Pal:       p,
		currentHp: p.BaseHp * 10,
	}
}

func (player *battlePal) Attack(opponent *battlePal) int {
	dmg := 100 + (player.BaseAttack - opponent.BaseDefence)
	attackType := player.Types[0]
	for _, playerType := range player.Types {
		scale := playerType.Stronger(opponent.Types)
		dmg *= scale
	}
	for _, opponentType := range opponent.Types {
		scale := opponentType.Stronger([]pal.Type{attackType})
		dmg /= scale
	}
	opponent.currentHp -= dmg
	return opponent.currentHp
}

// Duel battles 2 pals
//   - TODO: take defence into account
func (s *SimpleBattler) Duel(playerPal, opponentPal pal.Pal) Result {
	player := newBattlePal(playerPal)
	opponent := newBattlePal(opponentPal)
	log.Printf("%d vs %d", player.currentHp, opponent.currentHp)
	for player.currentHp > 0 && opponent.currentHp > 0 {
		if player.Speed == opponent.Speed {
			player.Attack(opponent)
			opponent.Attack(player)
			continue
		}
		if player.Speed > opponent.Speed {
			if player.Attack(opponent) <= 0 {
				return ResultWin
			}
		}
		if opponent.Attack(player) <= 0 {
			return ResultLoose
		}
		if opponent.Speed > player.Speed {
			if player.Attack(opponent) <= 0 {
				return ResultWin
			}
		}
		log.Printf("%d vs %d", player.currentHp, opponent.currentHp)
	}
	if player.currentHp > 0 {
		return ResultWin
	}
	if opponent.currentHp > 0 {
		return ResultLoose
	}
	return ResultDraw
}
