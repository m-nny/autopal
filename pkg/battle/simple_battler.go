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

func (player *battlePal) Attack(opponent *battlePal) {
	dmg := player.BaseAttack
	attackType := player.Types[0]
	for _, playerType := range player.Types {
		dmg *= playerType.Stronger(opponent.Types)
	}
	for _, opponentType := range opponent.Types {
		dmg /= opponentType.Stronger([]pal.Type{attackType})
	}
	opponent.currentHp -= dmg
}

// Duel battles 2 pals
//   - TODO: take defence into account
//   - TODO: take types into effect
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
			player.Attack(opponent)
			if opponent.currentHp <= 0 {
				return ResultWin
			}
		}
		opponent.Attack(player)
		if player.currentHp <= 0 {
			return ResultLoose
		}
		if opponent.Speed > player.Speed {
			player.Attack(opponent)
			if opponent.currentHp <= 0 {
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
