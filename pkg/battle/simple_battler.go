package battle

import (
	"log"

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
}

func NewSimpleBattler() *SimpleBattler {
	return &SimpleBattler{}
}

// Duel battles 2 pals
//   - TODO: take speed into account
//   - TODO: take defence into account
//   - TODO: take types into effect
func (s *SimpleBattler) Duel(player, opponent pal.Pal) Result {
	playerHp := player.BaseHp
	opponentHp := opponent.BaseHp
	log.Printf("%d vs %d", playerHp, opponentHp)
	for playerHp > 0 && opponentHp > 0 {
		playerHp -= opponent.BaseAttack
		opponentHp -= player.BaseAttack
		log.Printf("%d vs %d", playerHp, opponentHp)
	}
	if playerHp > 0 {
		return ResultWin
	}
	if opponentHp > 0 {
		return ResultLoose
	}
	return ResultDraw
}
